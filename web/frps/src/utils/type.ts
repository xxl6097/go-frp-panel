// 定义类型化的注入键
export interface Version {
  frpcVersion: string
  appName: string
  appVersion: string
  buildTime: string
  gitRevision: string
  gitBranch: string
  goVersion: string
  displayName: string
  description: string
  osType: string
  arch: string
  compiler: string
  gitTreeState: string
  gitCommit: string
  gitVersion: string
  gitReleaseCommit: string
  binName: string
  totalSize: string
  usedSize: string
  freeSize: string
}

export interface Client {
  osType: string
  secKey: string
  devMac: string
  devIp: string
  frpId: string
  sseId: string
  hostName: string
  appVersion: string
}

export interface Option {
  value: string
  label: string
}
