import { ElLoading, ElMessage, ElMessageBox } from 'element-plus'
import { EventAwareSSEClient } from "./sseclient.ts";

export function deepCopyJSON<T>(obj: T): T {
  return JSON.parse(JSON.stringify(obj))
}

export function markdownToHtml(markdown: string): string {
  const lines: string[] = markdown.split('\n')
  let html: string = ''
  let inList: boolean = false
  let listItems: string[] = []
  let inCodeBlock: boolean = false
  let codeBlockContent: string = ''

  for (let i = 0; i < lines.length; i++) {
    const line: string = lines[i].trim()

    // 处理代码块开始
    if (line.startsWith('```')) {
      if (inCodeBlock) {
        html += `<pre><code>${codeBlockContent}</code></pre>`
        inCodeBlock = false
        codeBlockContent = ''
      } else {
        inCodeBlock = true
      }
      continue
    }

    if (inCodeBlock) {
      codeBlockContent += line + '\n'
      continue
    }

    // 处理标题
    if (/^(#+) (.*)$/.test(line)) {
      const [, hashes, content] = line.match(/^(#+) (.*)$/)!
      const level: number = hashes.length
      if (inList) {
        html += `<ul>${listItems.join('')}</ul>`
        inList = false
        listItems = []
      }
      html += `<h${level}>${content}</h${level}>`
    }
    // 处理无序列表
    else if (/^([*-]) (.*)$/.test(line)) {
      const [, , content] = line.match(/^([*-]) (.*)$/)!
      if (!inList) {
        inList = true
      }
      listItems.push(`<li>${content}</li>`)
    }
    // 处理段落
    else {
      if (inList) {
        html += `<ul>${listItems.join('')}</ul>`
        inList = false
        listItems = []
      }
      if (line) {
        html += `<p>${line}</p>`
      }
    }
  }

  // 如果最后处于列表状态，闭合列表
  if (inList) {
    html += `<ul>${listItems.join('')}</ul>`
  }

  // 如果最后处于代码块状态，闭合代码块
  if (inCodeBlock) {
    html += `<pre><code>${codeBlockContent}</code></pre>`
  }

  // 处理加粗
  html = html.replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>')

  // 处理斜体
  html = html.replace(/\*(.*?)\*/g, '<em>$1</em>')

  return html
}

// function uint8ArrayToString(array: Uint8Array): string {
//   const decoder = new TextDecoder('utf-8') // 可替换为其他编码如 'ascii'
//   return decoder.decode(array)
// }

// function bytesToAsciiString(bytes: Uint8Array): string {
//   let str = ''
//   for (const byte of bytes) {
//     str += String.fromCharCode(byte)
//   }
//   return str
// }

function bytesToHexString(bytes: Uint8Array): string {
  return Array.from(bytes)
    .map((byte) => byte.toString(16).padStart(2, '0'))
    .join('')
}

// 生成 32 字节随机密钥（256 位）
function generateRandomKeys(length: number): string {
  const array = new Uint8Array(length)
  window.crypto.getRandomValues(array)
  const key = bytesToHexString(array)
  return key
}

export function generateRandomKey(): string {
  return generateRandomKeys(32)
}

// const key = generateRandomKey(32);

export function generateRandomKey1(length: number) {
  const characters =
    'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
  let key = ''
  for (let i = 0; i < length; i++) {
    const randomIndex = Math.floor(Math.random() * characters.length)
    key += characters.charAt(randomIndex)
  }
  return key
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

export function showInfoTips(message: string) {
  ElMessage({
    showClose: true,
    message: message,
    type: 'info',
  })
}

export const getFilePathByValue = (opt: any, valuePath: any) => {
  const child = opt.find((item: any) => item.value === valuePath[0])
  if (child) {
    const children = child.children
    if (children) {
      const node = children.find((item: any) => item.value === valuePath[1])
      if (node) {
        console.log(node)
        return node
      }
    }
  }
  return null
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

export function downloadByPost(title: string, path: string, body: any) {
  return new Promise((resolve, reject) => {
    let loading: any
    if (title !== undefined) {
      loading = showLoading(title)
    }

    const header = {
      'Content-Type': 'application/json',
    }
    fetch(path, {
      credentials: 'include',
      method: 'POST',
      headers: header,
      body: body,
    })
      .then(async (response) => {
        //if (!response.ok) throw new Error(`HTTP ${response.status}`)
        if (response.ok) {
          const disposition = response.headers.get('Content-Disposition')
          const filename = getFilenameFromContentDisposition(
            disposition as string,
          )
          return response.blob().then((blob) => ({ blob, filename }))
        } else {
          const text = await response.text()
          console.log('downloadByPost', text, response)
          showErrorTips(text)
          throw new Error(`HTTP ${response.statusText} ${text}`)
        }
      })
      .then(({ blob, filename }) => {
        const link = document.createElement('a')
        link.href = window.URL.createObjectURL(blob)
        link.download = filename as string
        link.style.display = 'none'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        resolve(filename)
      })
      .catch((error) => {
        console.log('downloadByPost', path, error)
        showErrorTips(error.message)
        reject(error.message)
      })
      .finally(() => {
        if (loading) {
          loading.close()
        }
      })
  })
}

export function DownLoadFile(
  title: string,
  method: string,
  path: string,
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
      body: body,
    })
      .then(async (response) => {
        //if (!response.ok) throw new Error(`HTTP ${response.status}`)
        if (response.ok) {
          const disposition = response.headers.get('Content-Disposition')
          const filename = getFilenameFromContentDisposition(
            disposition as string,
          )
          return response.blob().then((blob) => ({ blob, filename }))
        } else {
          const text = await response.text()
          console.log('DownLoadFile', text, response)
          showErrorTips(text)
          throw new Error(`HTTP ${response.statusText} ${text}`)
        }
      })
      .then(({ blob, filename }) => {
        const link = document.createElement('a')
        link.href = window.URL.createObjectURL(blob)
        link.download = filename as string
        link.style.display = 'none'
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        resolve(filename)
      })
      .catch((error) => {
        console.log('DownLoadFile', path, error)
        showErrorTips(error.message)
        reject(error.message)
      })
      .finally(() => {
        if (loading) {
          loading.close()
        }
      })
  })
}

export async function downloadByPost1(url: string, body: any) {
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
        if (json.code !== 0) {
          reject(json.msg)
          if (json.msg !== '') {
            showErrorTips(json.msg)
          }
        } else {
          resolve(json.data)
          if (json.msg !== '') {
            //showSucessTips(json.msg)
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
          //config.onUploadProgress(percentComplete.toFixed(2))
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

async function uploadChunk(
  api: string,
  method: string,
  chunk: any,
  fileName: string,
  chunkIndex: any,
  totalChunks: any,
  hash: any,
) {
  const response = await fetch(api, {
    method: method,
    headers: {
      'X-File-Name': fileName,
      'X-Chunk-Index': chunkIndex,
      'X-Total-Chunks': totalChunks,
      'X-File-Hash': hash,
    },
    body: chunk,
  })
  if (!response.ok) {
    throw new Error(`Upload failed with status ${response.status}`)
  }
}

export async function piecesUpload(
  api: string,
  method: string,
  file: any,
  onProgress: any,
  onFinish: any,
) {
  const chunkSize = 1024 * 1024 // 512KB
  const totalChunks = Math.ceil(file.size / chunkSize)
  const hash = Date.now().toString()

  let uploadedChunks = 0
  const concurrency = 50 // 最多并发 3 个请求

  // 分批次上传（使用队列控制并发）
  const queue: any[] = []
  for (let i = 0; i < totalChunks; i++) {
    queue.push(i)
  }

  async function processQueue() {
    const currentPromises = []
    while (queue.length > 0 && currentPromises.length < concurrency) {
      const chunkIndex = queue.shift()
      const start = chunkIndex * chunkSize
      const end = Math.min(start + chunkSize, file.size)
      const chunk = file.slice(start, end)

      const promise = uploadChunk(
        api,
        method,
        chunk,
        file.name,
        chunkIndex,
        totalChunks,
        hash,
      )
        .then(() => {
          uploadedChunks++
          const progress = (uploadedChunks / totalChunks) * 100
          if (onProgress) {
            onProgress(progress.toFixed(2))
          }
        })
        .catch((error) => {
          console.error(`Chunk ${chunkIndex} failed:`, error)
          queue.unshift(chunkIndex) // 失败的分片重新加入队列
        })

      currentPromises.push(promise)
    }

    await Promise.all(currentPromises)
    if (queue.length > 0) {
      await processQueue()
    }
  }

  await processQueue()
  if (onFinish) {
    onFinish()
  }
}

// function test(){
//   // 使用示例
//   xhrPromise({
//     url: 'https://api.example.com/data',
//     method: 'POST',
//     data: JSON.stringify({ key: 'value' }),
//     headers: {
//       'Content-Type': 'application/json'
//     },
//     onUploadProgress: ({ progress }) => {
//       console.log(`上传进度：${(progress * 100).toFixed(2)}%`);
//     },
//     onDownloadProgress: ({ progress }) => {
//       console.log(`下载进度：${(progress * 100).toFixed(2)}%`);
//     }
//   })
//       .then(response => {
//         console.log('请求成功', response);
//       })
//       .catch(error => {
//         console.error('请求失败', error);
//       });
// }

export function testSSEClient() {
  // 使用示例
  const client = new EventAwareSSEClient('http://api.example.com/sse')
  client.addEventListener('stockUpdate', (data) => {
    console.log('股票更新:', data.price)
  })
}