import {setEditorClasses} from '../utils';

export default class Paragraph {
  constructor({data, config, api, readOnly}) {
    this.api = api;
    this.readOnly = readOnly;
    this.config = config;
    this._data = {};
    this._element = this.drawView();
    this._preserveBlank = config.preserveBlank !== undefined ? config.preserveBlank : false;
    this.data = data;
  }

  onKeyUp(e) {
    if (e.code !== 'Backspace' && e.code !== 'Delete') {
      return;
    }
    const {textContent} = this._element;
    if (textContent === '') {
      this._element.innerHTML = '';
    }
  }

  drawView() {
    let el = document.createElement('p');
    setEditorClasses(this.config, el);
    el.contentEditable = false;
    if (!this.readOnly) {
      el.contentEditable = true;
      el.addEventListener('keyup', this.onKeyUp.bind(this));
    }
    return el;
  }

  render() {
    return this._element;
  }

  merge(data) {
    let newData = {
      text : this.data.text + data.text
    };

    this.data = newData;
  }

  validate(savedData) {
    if(savedData.text.trim() === '' && !this._preserveBlank) {
      return false;
    }
    return true;
  }

  save(toolsContent) {
    return {
      text: toolsContent.innerHTML
    };
  }

  static get conversionConfig() {
    return {
      export: 'text', // to convert Paragraph to other block, use 'text' property of saved data
      import: 'text' // to covert other block's exported string to Paragraph, fill 'text' property of tool data
    };
  }

  static get sanitize() {
    return {
      text: {
        br: true,
      }
    };
  }

  get data() {
    let text = this._element.innerHTML;
    this._data.text = text;
    return this._data;
  }
  set data(data) {
    this._data = data || {};
    this._element.innerHTML = this._data.text || '';
  }

  static get toolbox() {
    return {
      icon: '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="black" width="24px" height="24px"><path d="M0 0h24v24H0z" fill="none"/><path d="M5 17v2h14v-2H5zm4.5-4.2h5l.9 2.2h2.1L12.75 4h-1.5L6.5 15h2.1l.9-2.2zM12 5.98L13.87 11h-3.74L12 5.98z"/></svg>',
      title: 'Text'
    };
  }
}