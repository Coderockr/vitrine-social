import axios from 'axios';

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

const api = ax();

export default api;
