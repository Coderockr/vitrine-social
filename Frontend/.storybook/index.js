import '../src/App.css';
import insertCss from 'insert-css';
import css from 're-bulma/build/css';

try {
  if (typeof document !== 'undefined' || document !== null) insertCss(css, { prepend: true });
} catch (e) {
  console.log(e)
}

import ClassifiedCard from '../src/components/ClassifiedCard/story';
