(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[888],{3454:function(e,t,n){"use strict";var r,o;e.exports=(null===(r=n.g.process)||void 0===r?void 0:r.env)&&"object"===typeof(null===(o=n.g.process)||void 0===o?void 0:o.env)?n.g.process:n(7663)},1118:function(e,t,n){(window.__NEXT_P=window.__NEXT_P||[]).push(["/_app",function(){return n(8484)}])},5463:function(e,t,n){"use strict";var r=(0,n(7294).createContext)();t.Z=r},8484:function(e,t,n){"use strict";n.r(t);var r=n(5893),o=n(7294),u=n(5463),i=n(1163),c=n(3454);function s(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function f(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{},r=Object.keys(n);"function"===typeof Object.getOwnPropertySymbols&&(r=r.concat(Object.getOwnPropertySymbols(n).filter((function(e){return Object.getOwnPropertyDescriptor(n,e).enumerable})))),r.forEach((function(t){s(e,t,n[t])}))}return e}t.default=function(e){var t=e.Component,n=e.pageProps,s=(0,o.useState)(c.env.APP_KEY)[0],a=(0,o.useState)(""),l=a[0],p=a[1],h=(0,o.useState)(""),v=h[0],d=h[1],y=(0,i.useRouter)();return(0,o.useEffect)((function(){y.push("/")}),[]),(0,r.jsx)(u.Z.Provider,{value:{appKey:s,clientKey:l,setClientKey:p,userToken:v,setUserToken:d},children:(0,r.jsx)(t,f({},n))})}},7663:function(e){!function(){var t={162:function(e){var t,n,r=e.exports={};function o(){throw new Error("setTimeout has not been defined")}function u(){throw new Error("clearTimeout has not been defined")}function i(e){if(t===setTimeout)return setTimeout(e,0);if((t===o||!t)&&setTimeout)return t=setTimeout,setTimeout(e,0);try{return t(e,0)}catch(r){try{return t.call(null,e,0)}catch(r){return t.call(this,e,0)}}}!function(){try{t="function"===typeof setTimeout?setTimeout:o}catch(e){t=o}try{n="function"===typeof clearTimeout?clearTimeout:u}catch(e){n=u}}();var c,s=[],f=!1,a=-1;function l(){f&&c&&(f=!1,c.length?s=c.concat(s):a=-1,s.length&&p())}function p(){if(!f){var e=i(l);f=!0;for(var t=s.length;t;){for(c=s,s=[];++a<t;)c&&c[a].run();a=-1,t=s.length}c=null,f=!1,function(e){if(n===clearTimeout)return clearTimeout(e);if((n===u||!n)&&clearTimeout)return n=clearTimeout,clearTimeout(e);try{n(e)}catch(t){try{return n.call(null,e)}catch(t){return n.call(this,e)}}}(e)}}function h(e,t){this.fun=e,this.array=t}function v(){}r.nextTick=function(e){var t=new Array(arguments.length-1);if(arguments.length>1)for(var n=1;n<arguments.length;n++)t[n-1]=arguments[n];s.push(new h(e,t)),1!==s.length||f||i(p)},h.prototype.run=function(){this.fun.apply(null,this.array)},r.title="browser",r.browser=!0,r.env={},r.argv=[],r.version="",r.versions={},r.on=v,r.addListener=v,r.once=v,r.off=v,r.removeListener=v,r.removeAllListeners=v,r.emit=v,r.prependListener=v,r.prependOnceListener=v,r.listeners=function(e){return[]},r.binding=function(e){throw new Error("process.binding is not supported")},r.cwd=function(){return"/"},r.chdir=function(e){throw new Error("process.chdir is not supported")},r.umask=function(){return 0}}},n={};function r(e){var o=n[e];if(void 0!==o)return o.exports;var u=n[e]={exports:{}},i=!0;try{t[e](u,u.exports,r),i=!1}finally{i&&delete n[e]}return u.exports}r.ab="//";var o=r(162);e.exports=o}()},1163:function(e,t,n){e.exports=n(880)}},function(e){var t=function(t){return e(e.s=t)};e.O(0,[774,179],(function(){return t(1118),t(880)}));var n=e.O();_N_E=n}]);