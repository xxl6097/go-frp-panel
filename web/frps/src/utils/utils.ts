import {ElLoading, ElMessage, ElMessageBox} from 'element-plus'
export function deepCopyJSON<T>(obj: T): T {
  return JSON.parse(JSON.stringify(obj));
}
export function generateRandomKey(length:number) {
  const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
  let key = '';
  for (let i = 0; i < length; i++) {
    const randomIndex = Math.floor(Math.random() * characters.length);
    key += characters.charAt(randomIndex);
  }
  return key;
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

export function showWarmDialog(title:string,ok:any,cancel:any) {
  ElMessageBox.confirm(
      title,
      'Warning',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }
  )
      .then(() => {
        ok()
      })
      .catch(() => {
        cancel()
      })
}

export function downloadFile(url:string) {
  fetch(url, {
    method: 'GET',
  })
      .then(response => {
        // 获取 Content-Disposition 头信息
        const disposition = response.headers.get('Content-Disposition');
        let filename = 'downloaded_file'; // 默认文件名

        if (disposition && disposition.includes('filename=')) {
          const matches = disposition.match(/filename\*=UTF-8''(.+)|filename="?(.+?)"?$/);
          if (matches) {
            filename = decodeURIComponent(matches[1] || matches[2]);
          }
        }

        return response.blob().then(blob => ({ blob, filename }));
      })
      .then(({ blob, filename }) => {
        // 创建下载链接
        const link = document.createElement('a');
        link.href = URL.createObjectURL(blob);
        link.download = filename;
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
      })
      .catch(error => console.error('下载失败:', error));
}

export function getFilenameFromContentDisposition(contentDisposition:string) {
  if (!contentDisposition) return null;
  const matches = contentDisposition.match(/filename="?([^"]+)"?/);
  return matches && matches[1]? matches[1] : null;
}


export function downloadByPost(title:string,path:string,body:any){
  return new Promise((resolve, reject)=>{
    let loading:any;
    if (title !== undefined){
      loading = showLoading(title)
    }

    const header = {
      'Content-Type': 'application/json'
    }
    fetch(path, {
      credentials: 'include',
      method: 'POST',
      headers: header,
      body: body,
    })
        .then((response) => {
          if (!response.ok) throw new Error(`HTTP ${response.status}`);
          const disposition = response.headers.get('Content-Disposition');
          const filename = getFilenameFromContentDisposition(disposition as string)
          return response.blob().then(blob => ({ blob, filename }));
        })
        .then(({ blob, filename }) => {
          const link = document.createElement('a');
          link.href = window.URL.createObjectURL(blob);
          link.download = filename as string;
          link.style.display = 'none';
          document.body.appendChild(link);
          link.click();
          document.body.removeChild(link);
          resolve(filename)
    })
        .catch((error) => {
          console.log(path,error)
          reject(error.message);
          showErrorTips(error.message)
        })
        .finally(() => {
          if (loading){
            loading.close()
          }
        })
  })
}


export async function downloadByPost1(url:string,body:any){
  try {
    const header = {
      'Content-Type': 'application/json'
    }
    const response = await fetch(url,
        {
          method: 'POST',
          credentials: 'include',
          body: body,
          headers: header,
        });
    if (!response.ok) throw new Error(`HTTP ${response.status}`);
    const disposition = response.headers.get('Content-Disposition');
    const filename = getFilenameFromContentDisposition(disposition as string)
    const blob = await response.blob();
    const link = document.createElement('a');
    link.href = window.URL.createObjectURL(blob);
    link.download = filename as string;
    link.style.display = 'none';
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);

  } catch (error:any) {
    throw new Error(`文件下载失败: ${error.message}`);
  }
}


export async function download(url:string){
  try {
    const response = await fetch(url, { method: 'GET', credentials: 'include'});
    if (!response.ok) throw new Error(`HTTP ${response.status}`);


    // const blob = await response.blob();
    // const link = document.createElement('a');
    // link.href = URL.createObjectURL(blob);
    // link.download = filename;
    // link.click();


    const disposition = response.headers.get('Content-Disposition');
    const filename = getFilenameFromContentDisposition(disposition as string)
    const blob = await response.blob();
    const link = document.createElement('a');
    link.href = window.URL.createObjectURL(blob);
    link.download = filename as string;
    link.style.display = 'none';
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);

  } catch (error:any) {
    throw new Error(`文件下载失败: ${error.message}`);
  }
}

export function post(title:string,path:string,body:any){
  return fetchReest('POST',title,path,body)
}
export function put(title:string,path:string,body:any){
  return fetchReest('PUT',title,path,body)
}
export function get(title:string,path:string,body:any){
  return fetchReest('GET',title,path,body)
}

export function fetchReest(method:string,title:string,path:string,body:any){
  const header = {
    'Content-Type': 'application/json'
  }
  return request(method,title,path,header,body)
}

export function request(method:string,title:string,path:string,header:any,body:any){
  return new Promise((resolve, reject)=>{
    let loading:any;
    if (title !== undefined){
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
          if (json.code !== 0){
            reject(json.msg)
            if (json.msg !== ''){
              showErrorTips(json.msg)
            }
          }else {
            resolve(json.data)
            if (json.msg !== ''){
              //showSucessTips(json.msg)
            }
          }
        })
        .catch((error) => {
          console.log(method,path,error)
          reject(error.message);
          showErrorTips(error.message)
        })
        .finally(() => {
          if (loading){
            loading.close()
          }
        })
  })
}
