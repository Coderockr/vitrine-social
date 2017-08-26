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

import ClassifiedCard from '../src/components/ClassifiedCard/story';
