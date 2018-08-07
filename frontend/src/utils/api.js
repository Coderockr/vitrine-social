import axios from 'axios';
import { deauthorizeUser } from './auth';
import BottomNotification from '../components/BottomNotification';

const storage = window.localStorage;

export const headers = { 'Content-type': 'application/x-www-form-urlencoded' };

const setHeaders = () => {
  const headerObj = { ...headers };
  const token = storage.getItem('token');
  if (token) {
    headerObj.Authorization = token;
  }
  return headerObj;
};

const ax = () => axios.create({
  baseURL: process.env.REACT_APP_API_URL,
  timeout: 5000,
  headers: setHeaders(),
});

const imageax = () => axios.create({
  baseURL: process.env.REACT_APP_API_URL,
  timeout: 50000,
  headers: setHeaders(),
});

const api = ax();
const apiImage = imageax();

api.interceptors.response.use(response => response,
  (error) => {
    if (error.response && error.response.data) {
      const { code, message } = error.response.data;
      if (code === 401 && message === 'Token Expired, get a new one') {
        return BottomNotification({
          message: 'Sua sessão expirou! Faça login novamente!',
          success: false,
          onClose: () => {
            deauthorizeUser();
            window.location.replace(`${process.env.REACT_APP_HOST}login`);
          },
        });
      }
    }
    return Promise.reject(error);
  });

export { api, apiImage };
