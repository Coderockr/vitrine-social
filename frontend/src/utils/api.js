import axios from 'axios';

export const headers = { 'Content-type': 'application/x-www-form-urlencoded' };

const ax = () => axios.create({
  baseURL: process.env.REACT_APP_API_URL,
  timeout: 5000,
});

const api = ax();

export default api;
