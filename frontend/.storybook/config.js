import { configure } from '@kadira/storybook';

function loadStories() {
  require('./index');
}

configure(loadStories, module);
