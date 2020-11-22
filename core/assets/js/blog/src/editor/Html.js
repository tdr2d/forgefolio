import {setEditorClasses} from '../utils';
import CodeFlask from 'codeflask';

const LineHeight = 20;
const ScriptURLRegex = new RegExp(/<script.+src="(https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*))".*><\/script>/gi)


export default class Html {
    static get toolbox() {
      return {
        title: 'Embed Html code',
        icon: '<svg xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 0 24 24" width="24"><path d="M0 0h24v24H0V0z" fill="none"/><path d="M9.4 16.6L4.8 12l4.6-4.6L8 6l-6 6 6 6 1.4-1.4zm5.2 0l4.6-4.6-4.6-4.6L16 6l6 6-6 6-1.4-1.4z"/></svg>'
      };
    }

    constructor({data, config, api, readOnly}) {
        this.data = data;
        this.config = config;
        this.api = api;
        this.readOnly = readOnly;
        this.blockSettings = [
          {
            name: 'code',
            title: 'Code View',
            icon: Html.toolbox.icon,
            el: document.createElement('button')
          },
          {
            name: 'visual',
            title: 'Visual View',
            icon: `<svg xmlns="http://www.w3.org/2000/svg" enable-background="new 0 0 24 24" height="24" viewBox="0 0 24 24" width="24"><g><rect fill="none" height="24" width="24"/><path d="M19,3H5C3.89,3,3,3.9,3,5v14c0,1.1,0.89,2,2,2h14c1.1,0,2-0.9,2-2V5C21,3.9,20.11,3,19,3z M19,19H5V7h14V19z M13.5,13 c0,0.83-0.67,1.5-1.5,1.5s-1.5-0.67-1.5-1.5c0-0.83,0.67-1.5,1.5-1.5S13.5,12.17,13.5,13z M12,9c-2.73,0-5.06,1.66-6,4 c0.94,2.34,3.27,4,6,4s5.06-1.66,6-4C17.06,10.66,14.73,9,12,9z M12,15.5c-1.38,0-2.5-1.12-2.5-2.5c0-1.38,1.12-2.5,2.5-2.5 c1.38,0,2.5,1.12,2.5,2.5C14.5,14.38,13.38,15.5,12,15.5z"/></g></svg>`,
            el: document.createElement('button')
          }
        ];
        this._html = this.data && this.data.html ? atob(this.data.html) : '',
        this.wrapper = this._getDefaultWrapper();
        this._view = this.data && this.data.view ? this.data.view : 'code';
        if (this._view == 'visual') {
          this._setVisualView();
        } else if (this._view == 'code') {
          this._setCodeView();
        }
    }
    
    _setCodeView() {
      this.wrapper.innerHTML = "";
      this.wrapper.className = "";
      this.flask = new CodeFlask(this.wrapper, { language: 'html' });
      this.flask.updateCode(this._html);
      this.flask.onUpdate((code) => { 
        this._html = code; 
        const height = LineHeight * code.split('\n').length;
        this.wrapper.style.minHeight = `${height}px`;
      });
    }

    _setVisualView() {
      this._loadScripts();
      setEditorClasses(this.config, this.wrapper);
      this.wrapper.innerHTML = "";
      this.wrapper.style.minHeight = "25px";
      this.wrapper.innerHTML = this._html;
    }

    _loadScripts() {
      try {
        const scripts = [...this._html.matchAll(ScriptURLRegex)];
        for (const i  in scripts) {
          const url = scripts[i][1];
          console.warn(`Loading script ${url}`);
          let script = document.createElement('script');
          script.setAttribute('src', url);
          script.async = true;
          document.head.appendChild(script);
        }
      } catch (e) {
        console.error(e);
      }
    }

    _getDefaultWrapper() {
      const wrapper = document.createElement('div');
      wrapper.style.minHeight = "25px";
      return wrapper;
    }

    renderSettings(){
      const wrapper = document.createElement('div');
      this.blockSettings.forEach( tune => {
        tune.el.innerHTML = tune.icon;
        tune.el.classList.add('cdx-settings-button');
        this.api.tooltip.onHover(tune.el, tune.title);
        wrapper.appendChild(tune.el);
        if (tune.name == "code") {
          if (this._view == "code") {tune.el.classList.add('cdx-settings-button--active');}
          tune.el.addEventListener('click', () => {
            this._setCodeView();
            tune.el.classList.add('cdx-settings-button--active');
            this.blockSettings[1].el.classList.remove('cdx-settings-button--active');
          });
        } else if (tune.name == "visual") {
          if (this._view == "visual") {tune.el.classList.add('cdx-settings-button--active');}
          tune.el.addEventListener('click', () => {
            this._setVisualView();
            tune.el.classList.add('cdx-settings-button--active');
            this.blockSettings[0].el.classList.remove('cdx-settings-button--active');
          });
        }
      });
      return wrapper;
    }

    render(){
      return this.wrapper;
    }
    
    save(){
      return {
        html: btoa(this._html),
        view: this._view
      };
    }

    validate(data) {
      return data.html && data.html.trim() !== '';
    }
  }