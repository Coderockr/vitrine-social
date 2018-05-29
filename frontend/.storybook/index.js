import '../src/App.css';
import insertCss from 'insert-css';
import css from 're-bulma/build/css';
import { updateLocale } from 'moment';
import ptBr from 'moment/locale/pt-br';

updateLocale('pt-br', ptBr)

try {
  if (typeof document !== 'undefined' || document !== null) insertCss(css, { prepend: true });
} catch (e) {
  console.log(e)
}

import App from '../src/containers/App/story';
import Categories from '../src/components/Categories/story';
import Dialog from '../src/components/Dialog/story';
import Pagination from '../src/components/Pagination/story';
import RequestCard from '../src/components/RequestCard/story';
import Requests from '../src/components/Requests/story';
import Header from '../src/components/Header/story';
import Search from '../src/components/Search/story';
