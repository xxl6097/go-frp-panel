export interface NetWork {
  name: string
  displayName: string
  macAddress: string
  ipv4: string
  ipAddresses: string[]
}

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
  hostName: string
  network: NetWork
}

export interface FrpcConfiguration {
  user: string
  token: string
  comment: string
  ports: string
  domains: string
  subdomains: string
  enable: boolean
  editable: boolean
  count: number
  id: string
}

export interface Client {
  osType: string
  secKey: string
  devMac: string
  devIp: string
  frpId: string
  sseId: string
  devName: string
  appVersion: string
}

export interface Option {
  value: string
  label: string
}

// {
//   "localIP": "0.0.0.0",
//   "localPort": 6401,
//   "name": "frpc.ui",
//   "remotePort": 6109,
//   "type": "tcp"
// }
export interface TypedProxyConfig {
  localIP: string
  localPort: number
  name: string
  remotePort: number
  type: string
}

export interface WebServerConfig {
  addr: string
  port: number
  user: string
  password: string
}

export interface ClientConfig {
  serverAddr: string
  serverPort: number
  proxies: Partial<TypedProxyConfig[]>
  webServer: Partial<WebServerConfig>
  metadatas: any
}

export interface UserConfig {
  id: string
  user: string
  token: string
  comment: string
  ports: any[]
  domains: string[]
  subdomains: string[]
  enable: boolean
}

export interface ConfigBodyData {
  binAddress: string
  serverAdminPort: number | undefined
  userConfig: Partial<UserConfig>
  clientConfig: Partial<ClientConfig>
}
