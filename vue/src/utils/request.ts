/*
 * @Author: Callay 2415993100@qq.com
 * @Date: 2024-05-06 23:25:30
 * @LastEditors: Callay 2415993100@qq.com
 * @LastEditTime: 2024-05-07 15:01:02
 * @FilePath: \checkIn\vue\src\utils\request.ts
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
import axios, { type AxiosInstance, type AxiosRequestConfig } from "axios";

class HRequest {
  instance: AxiosInstance;

  constructor() {
    axios.defaults.headers["Content-Type"] = "application/json;charset=utf-8";
    this.instance = axios.create({
      baseURL: "/api/",
      timeout: 30000,
    });
    this.instance.interceptors.request.use(
      (config) => {
        config.headers['jwt-token']=sessionStorage.getItem("jwt-token")
        return config;
      },
      (error) => {
        console.error("request error:" + error);
        return error;
      }
    );
    this.instance.interceptors.response.use(
      (response) => {
        let res = response.data;

        if (typeof res === "string") {
          res = res ? JSON.parse(res) : res;
        }
        return res;
      },
      (error) => {
        console.error("respone error:" + error);
        return error;
      }
    );
  }

  get<T = any>(config: AxiosRequestConfig ): Promise<T> {
    return this.instance.request({ ...config, method: 'GET' });
  }

  post<T = any>(config: AxiosRequestConfig ): Promise<T> {
    return this.instance.request({ ...config, method: 'POST' });
  }

  delete<T = any>(config: AxiosRequestConfig ): Promise<T> {
    return this.instance.request({ ...config, method: 'DELETE' });
  }

  patch<T = any>(config: AxiosRequestConfig ): Promise<T> {
    return this.instance.request({ ...config, method: 'PATCH' });
  }
}

export default HRequest;
