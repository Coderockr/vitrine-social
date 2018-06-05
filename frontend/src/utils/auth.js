import api, { headers } from './api';

const storage = window.localStorage;

const authorizeAPI = (token) => {
  api.defaults.headers.common.Authorization = token;
  storage.setItem('token', token);
};

export const authorizeUser = (response) => {
  authorizeAPI(response.token);
  storage.setItem('loggedUser', JSON.stringify(response.organization));
};

export const getUser = () => (
  JSON.parse(storage.getItem('loggedUser'))
);

export const deauthorizeUser = () => {
  api.defaults.headers = headers;
  return storage.clear();
};
