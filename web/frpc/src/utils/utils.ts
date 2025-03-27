import { ElLoading, ElMessage, ElMessageBox } from 'element-plus'
export function deepCopyJSON<T>(obj: T): T {
  return JSON.parse(JSON.stringify(obj))
}
export function getTimestamp() {
  return new Date().getTime() // 输出：1632994993000
}
export function generateRandomKey(length: number) {
  const characters =
    'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
  let key = ''
  for (let i = 0; i < length; i++) {
    const randomIndex = Math.floor(Math.random() * characters.length)
    key += characters.charAt(randomIndex)
  }
  return key
}

export function getProxyName(prefix: string): string {
  return `${prefix}_${generateRandomKey(4)}`
}

export function showWarmTips(message: string) {
  ElMessage({
    showClose: true,
    message: message,
    type: 'warning',
  })
}
export function showErrorTips(message: string) {
  ElMessage({
    showClose: true,
    message: message,
    type: 'error',
  })
}

export function markdownToHtml(mdText: string) {
  const lines = mdText.split('\n')
  let html = ''
  let inCodeBlock = false
  let inList = false

  for (const line of lines) {
    // 代码块处理
    if (line.trim().startsWith('```')) {
      inCodeBlock = !inCodeBlock
      html += inCodeBlock ? '<pre><code>' : '</code></pre>\n'
      continue
    }
    if (inCodeBlock) {
      html += line + '\n'
      continue
    }

    // 各语法规则处理
    let processed = line
    processed = processed
      .replace(
        /^(#{1,6})\s+(.+)/,
        (_, hashes, content) =>
          `<h${hashes.length}>${content}</h${hashes.length}>`,
      )
      .replace(/\*\*(.+?)\*\*/g, '<strong>$1</strong>')
      .replace(/\*(.+?)\*/g, '<em>$1</em>')
      .replace(/^-\s+(.+)/, (_, content) => {
        if (!inList) {
          inList = true
          return '<ul>\n<li>' + content + '</li>'
        }
        return '<li>' + content + '</li>'
      })
      .replace(/^-{3,}/, '<hr>')

    // 列表闭合
    if (inList && !processed.startsWith('<li>')) {
      processed = '</ul>\n' + processed
      inList = false
    }

    // 段落处理
    if (processed.trim()) {
      html += processed.replace(/\n/g, '<br>') + '\n'
    }
  }
  return html
}
export function showMessageDialog(
  title: string,
  confirmButtonText: string,
  message: string,
) {
  return ElMessageBox.confirm(markdownToHtml(message), title, {
    confirmButtonText: confirmButtonText,
    cancelButtonText: '取消',
    dangerouslyUseHTMLString: true,
  })
}

export function showInfoTips(message: string) {
  ElMessage({
    showClose: true,
    message: message,
    type: 'info',
  })
}

export function showTips(code: any, message: string) {
  if (code === 0) {
    showSucessTips(message)
  } else {
    showWarmTips(message)
  }
}

export function showSucessTips(message: string) {
  ElMessage({
    showClose: true,
    message: message,
    type: 'success',
  })
}

export function showLoading(title: string) {
  return ElLoading.service({
    lock: true,
    text: title,
    background: 'rgba(0, 0, 0, 0.7)',
  })
}

export function showWarmDialog(title: string, ok: any, cancel: any) {
  ElMessageBox.confirm(title, 'Warning', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(() => {
      ok()
    })
    .catch(() => {
      cancel()
    })
}

export function downloadFile(url: string) {
  fetch(url, {
    method: 'GET',
    headers: {
      // 如果服务器需要鉴权，可以在这里添加 Authorization
    },
  })
    .then((response) => {
      // 获取 Content-Disposition 头信息
      const disposition = response.headers.get('Content-Disposition')
      let filename = 'downloaded_file' // 默认文件名

      if (disposition && disposition.includes('filename=')) {
        const matches = disposition.match(
          /filename\*=UTF-8''(.+)|filename="?(.+?)"?$/,
        )
        if (matches) {
          filename = decodeURIComponent(matches[1] || matches[2])
        }
      }

      return response.blob().then((blob) => ({ blob, filename }))
    })
    .then(({ blob, filename }) => {
      // 创建下载链接
      const link = document.createElement('a')
      link.href = URL.createObjectURL(blob)
      link.download = filename
      document.body.appendChild(link)
      link.click()
      document.body.removeChild(link)
    })
    .catch((error) => console.error('下载失败:', error))
}

export function getFilenameFromContentDisposition(contentDisposition: string) {
  if (!contentDisposition) return null
  const matches = contentDisposition.match(/filename="?([^"]+)"?/)
  return matches && matches[1] ? matches[1] : null
}

export async function downloadByPost(url: string, body: any) {
  try {
    const header = {
      'Content-Type': 'application/json',
    }
    const response = await fetch(url, {
      method: 'POST',
      credentials: 'include',
      body: body,
      headers: header,
    })
    if (!response.ok) throw new Error(`HTTP ${response.status}`)
    const disposition = response.headers.get('Content-Disposition')
    const filename = getFilenameFromContentDisposition(disposition as string)
    const blob = await response.blob()
    const link = document.createElement('a')
    link.href = window.URL.createObjectURL(blob)
    link.download = filename as string
    link.style.display = 'none'
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
  } catch (error: any) {
    throw new Error(`文件下载失败: ${error.message}`)
  }
}

export async function download(url: string) {
  try {
    const response = await fetch(url, { method: 'GET', credentials: 'include' })
    if (!response.ok) throw new Error(`HTTP ${response.status}`)

    // const blob = await response.blob();
    // const link = document.createElement('a');
    // link.href = URL.createObjectURL(blob);
    // link.download = filename;
    // link.click();

    const disposition = response.headers.get('Content-Disposition')
    const filename = getFilenameFromContentDisposition(disposition as string)
    const blob = await response.blob()
    const link = document.createElement('a')
    link.href = window.URL.createObjectURL(blob)
    link.download = filename as string
    link.style.display = 'none'
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
  } catch (error: any) {
    throw new Error(`文件下载失败: ${error.message}`)
  }
}

export function post(title: string, path: string, body: any) {
  return fetchReest('POST', title, path, body)
}
export function put(title: string, path: string, body: any) {
  return fetchReest('PUT', title, path, body)
}
export function get(title: string, path: string, body: any) {
  return fetchReest('GET', title, path, body)
}

export function fetchReest(
  method: string,
  title: string,
  path: string,
  body: any,
) {
  const header = {
    'Content-Type': 'application/json',
  }
  return request(method, title, path, header, body)
}

export function request(
  method: string,
  title: string,
  path: string,
  header: any,
  body: any,
) {
  return new Promise((resolve, reject) => {
    let loading: any
    if (title !== undefined) {
      loading = showLoading(title)
    }
    fetch(path, {
      credentials: 'include',
      method: method,
      headers: header,
      body: body,
    })
      .then((res) => {
        return res.json()
      })
      .then((json) => {
        resolve(json)
        if (json.code !== 0) {
          reject(json.msg)
          if (json.msg !== '') {
            showErrorTips(json.msg)
          }
        } else {
          if (json.msg !== '') {
            showSucessTips(json.msg)
          }
        }
      })
      .catch((error) => {
        console.log(method, path, error)
        reject(error.message)
        showErrorTips(error.message)
      })
      .finally(() => {
        if (loading) {
          loading.close()
        }
      })
  })
}

/**
 * 基于 Promise 封装的 XMLHttpRequest 请求
 * @param {Object} config - 请求配置
 * @param {string} config.url - 请求地址
 * @param {string} [config.method='GET'] - 请求方法
 * @param {Object} [config.headers] - 请求头
 * @param {any} [config.data] - 请求数据
 * @param {number} [config.timeout=0] - 超时时间（毫秒）
 * @param {string} [config.responseType] - 响应类型
 * @param {Function} [config.onUploadProgress] - 上传进度回调
 * @param {Function} [config.onDownloadProgress] - 下载进度回调
 * @returns {Promise} 返回 Promise 对象
 */
export function xhrPromise(config: any) {
  return new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest()
    // 初始化请求
    xhr.open(config.method || 'GET', config.url)
    // 设置请求头
    if (config.headers) {
      Object.entries(config.headers).forEach(([key, value]) => {
        xhr.setRequestHeader(key, value as string)
      })
    }

    // 设置响应类型
    if (config.responseType) {
      xhr.responseType = config.responseType
    }

    // 设置超时
    if (config.timeout) {
      xhr.timeout = config.timeout
    }

    // 上传进度处理
    if (config.onUploadProgress) {
      xhr.upload.onprogress = (event) => {
        if (event.lengthComputable) {
          const percentComplete = (event.loaded / event.total) * 100
          console.log('--->', percentComplete + '%')
          config.onUploadProgress(percentComplete.toFixed(2))
        }
      }
    }

    // 下载进度处理
    if (config.onDownloadProgress) {
      xhr.onprogress = (e) => {
        config.onDownloadProgress({
          loaded: e.loaded,
          total: e.total,
          progress: e.loaded / e.total,
        })
      }
    }

    // 请求成功处理
    xhr.onload = () => {
      if (xhr.status >= 200 && xhr.status < 300) {
        resolve({
          data: xhr.response,
          status: xhr.status,
          statusText: xhr.statusText,
          headers: xhr.getAllResponseHeaders(),
        })
      } else {
        reject(new Error(`请求失败：${xhr.status} ${xhr.statusText}`))
      }
    }

    // 错误处理
    xhr.onerror = () => reject(new Error('网络错误'))
    xhr.ontimeout = () => reject(new Error(`请求超时（${config.timeout}ms）`))
    xhr.onabort = () => reject(new Error('请求被中止'))

    // 发送请求
    try {
      xhr.send(config.data)
    } catch (err) {
      reject(err)
    }
  })
}
