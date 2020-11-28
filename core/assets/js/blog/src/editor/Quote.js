import {CreateElement} from '../utils';

export default class Quote {
  static get toolbox() {
    return {
      icon: feather.icons["message-circle"].toSvg(),
      title: 'Quote'
    };
  }

  constructor({data, config, api, readOnly}) {
    this.api = api;
    this.config = config;
    this.data = {
      text: data && data.text ? data.text : 'I love to travel, but I hate to arrive. (A. Einstein)'
    }
    this.nodes = {
      wrapper: CreateElement('blockquote', this.config.css && this.config.css.quote ? this.config.css.quote : []),
      paragraph: CreateElement('p', this.config.css && this.config.css.p ? this.config.css.p : [], {contentEditable: true, innerHTML: this.data.text})
    }
    this.nodes.paragraph.addEventListener('keyup', this.onKeyUp.bind(this));
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

  render() {
    this.nodes.wrapper.appendChild(this.nodes.paragraph);
    return this.nodes.wrapper;
  }

  save() {
    return {
      text: this.nodes.paragraph.innerHTML
    };
  }

  static get sanitize() {
    return {
      text: {
        br: true,
      }
    };
  }
}