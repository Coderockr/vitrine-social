import axios from 'axios';

export const headers = () => {
  const token = window.localStorage.getItem('token');
  if (token) {
    return {
      'Content-type': 'application/x-www-form-urlencoded',
      Authorization: token,
    };
  }
  return { 'Content-type': 'application/x-www-form-urlencoded' };
};

const api = axios.create({
  baseURL: process.env.REACT_APP_API_URL,
  timeout: 5000,
  headers: headers(),
});

export default api;
