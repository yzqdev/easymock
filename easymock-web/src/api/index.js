import axios from 'axios'
import {ElMessage} from 'element-plus'
import conf from '@/config/config'
let router
const instance = axios.create({
  baseURL:   `http://localhost:3901/api`,

})

const loading = {
  count: 0,
  isLoading: false,
  start () {
    this.count += 1
    if (!this.isLoading) {
      setTimeout(() => {
        if (!this.isLoading && this.count > 0) {
          this.isLoading = true
          this.checkLoading()
        }
      }, 1000)
    }
  },
  cancel () {
    this.count -= 1
    if (this.count <= 0) {
      this.done()
    }
  },
  done () {
    this.count = 0
    this.isLoading = false
    // iView.LoadingBar.finish()
  },
  checkLoading () {
   if (process.client) {
     const el = document.querySelector('.ivu-loading-bar')
     if (this.isLoading && !el) {
       // iView.LoadingBar.start()
     }
   }
  }
}

instance.interceptors.request.use((config) => {
  let token

    token = localStorage.getItem(conf.storageNamespace + 'token')

  config.headers.Authorization = `Bearer ${token}`
  return config
}, error => Promise.reject(error))

instance.interceptors.response.use((res) => {
  const messageUnless = res.config.messageUnless || []
  const body = res.data

  if (body.success === false) {
    if (body.code === 10001) {
      body.data.forEach((date) => {
        ElMessage.error({
          title: 'Error',
          desc: date[Object.keys(date)[0]]
        })
      })
    } else if (messageUnless.indexOf(body.message) === -1) {
      ElMessage.error({
        title: 'Error',
        desc: body.message
      })
    }
    return Promise.reject(res)
  }
  return res
}, (error) => {
  const res = error.response
  if (res) {
    if (res.status === 401 && /authentication/i.test(res.data.error)) {
        router.push('/log-out')
    } else if (  res.data && res.data.error) {
      ElMessage.error({
        title: 'Error',
        desc: res.data.error
      })
    }
  }
  Promise.reject(error).then(r => {return error})
})

const initAPI = _router => (router = _router)
const createAPI = (url, method, config) => {
  config = config || {}
  return instance({
    url,
    method,
    ...config
  })
}

// 创建用于导出数据的表单
const createExportExcelForm = (url, data) => {
  if (process.client) {
    const form = document.createElement('form')

    form.method = 'POST'
    form.action = url

    if (Array.isArray(data)) {
      data.forEach((d) => {
        const input = document.createElement('input')
        input.name = 'ids[]'
        input.value = d
        form.appendChild(input)
      })
    } else {
      const input = document.createElement('input')
      input.name = 'project_id'
      input.value = data
      form.appendChild(input)
    }

    document.body.appendChild(form)
    form.submit()
    document.body.removeChild(form)
  }
}

const util = {
  wallpaper: config => createAPI('/wallpaper', 'get', config)
}

const u = {
  getList: config => createAPI('/u', 'get', config),
  login: config => createAPI('/u/login', 'post', config),
  register: config => createAPI('/u/register', 'post', config),
  update: config => createAPI('/u/update', 'post', config)
}

const project = {
  getList: config => createAPI('/project', 'get', config),
  copy: config => createAPI('/project/copy', 'post', config),
  create: config => createAPI('/project/create', 'post', config),
  update: config => createAPI('/project/update', 'post', config),
  updateSwagger: config => createAPI('/project/sync/swagger', 'post', config),
  updateWorkbench: config => createAPI('/project/update_workbench', 'post', config),
  delete: config => createAPI('/project/delete', 'post', config)
}

const mock = {
  getList: config => createAPI('/mock', 'get', config),
  create: config => createAPI('/mock/create', 'post', config),
  update: config => createAPI('/mock/update', 'post', config),
  delete: config => createAPI('/mock/delete', 'post', config),
  export: config => createExportExcelForm('/api/mock/export', config)
}

const group = {
  getList: config => createAPI('/group', 'get', config),
  join: config => createAPI('/group/join', 'post', config),
  create: config => createAPI('/group/create', 'post', config),
  update: config => createAPI('/group/update', 'post', config),
  delete: config => createAPI('/group/delete', 'post', config)
}

const dashboard = {
  getList: config => createAPI('/dashboard', 'get', config)
}

export {
  u,
  project,
  mock,
  util,
  group,
  dashboard,
  initAPI
}
