(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[459],{6429:function(e,n,t){(window.__NEXT_P=window.__NEXT_P||[]).push(["/login",function(){return t(4754)}])},4754:function(e,n,t){"use strict";t.r(n),t.d(n,{default:function(){return v}});var s=t(4051),r=t.n(s),a=t(5893),c=t(9008),i=t.n(c),o=t(4298),l=t.n(o),u=t(7294),d=t(1163),m=t(5463),h=t(1664),p=t.n(h);function f(e,n,t,s,r,a,c){try{var i=e[a](c),o=i.value}catch(l){return void t(l)}i.done?n(o):Promise.resolve(o).then(s,r)}var x=t(9669);function v(){var e=(0,u.useContext)(m.Z),n=(0,d.useRouter)(),t=(0,u.useState)(""),s=t[0],c=t[1],o=(0,u.useState)(""),h=o[0],v=o[1],j=(0,u.useState)(""),y=j[0],N=j[1],g=(0,u.useState)(!1),w=(g[0],g[1]),b=function(){var t,a=(t=r().mark((function t(a){var i,o;return r().wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return a.preventDefault(),w(!0),c(a.target.username.value),v(a.target.password.value),t.prev=4,i={username:s,password:h},t.next=8,x.post("http://localhost:8000/api/userdata",i);case 8:(o=t.sent).data.clientKey?(e.setClientKey(o.data.clientKey),n.push("/questions"),w(!1)):(N("Invalid Form"),w(!1)),t.next=16;break;case 12:t.prev=12,t.t0=t.catch(4),N("Form submission failed"),w(!1);case 16:case"end":return t.stop()}}),t,null,[[4,12]])})),function(){var e=this,n=arguments;return new Promise((function(s,r){var a=t.apply(e,n);function c(e){f(a,s,r,c,i,"next",e)}function i(e){f(a,s,r,c,i,"throw",e)}c(void 0)}))});return function(e){return a.apply(this,arguments)}}();return(0,a.jsxs)("div",{children:[(0,a.jsxs)(i(),{children:[(0,a.jsx)("title",{children:"Workspace"}),(0,a.jsx)("meta",{name:"description",content:"Generated by create next app"}),(0,a.jsx)("link",{rel:"icon",href:"/favicon.ico"}),(0,a.jsx)("link",{href:"https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css",rel:"stylesheet",integrity:"sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC",crossOrigin:"anonymous"}),(0,a.jsx)(l(),{src:"https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js",integrity:"sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM",crossOrigin:"anonymous"})]}),(0,a.jsxs)("div",{className:"container my-5 p-5",children:[(0,a.jsx)("h1",{className:"text-primary d-flex justify-content-center",children:"Login"}),(0,a.jsx)("form",{onSubmit:b,className:"my-5",children:(0,a.jsxs)("div",{className:"row g-3",children:[(0,a.jsx)("div",{className:"row g-3 d-flex justify-content-center",children:(0,a.jsxs)("div",{className:"col-5",children:[(0,a.jsx)("label",{className:"my-2",htmlFor:"ssn",children:"Username"}),(0,a.jsx)("input",{type:"text",className:"form-control",id:"username",name:"username",value:s,placeholder:"Username",onChange:function(e){return c(e.target.value)}})]})}),(0,a.jsx)("div",{className:"row g-3 d-flex justify-content-center",children:(0,a.jsxs)("div",{className:"col-5",children:[(0,a.jsx)("label",{className:"my-2",htmlFor:"password",children:"Password"}),(0,a.jsx)("input",{type:"text",className:"form-control",id:"password1",name:"password",placeholder:"",value:h,onChange:function(e){return v(e.target.value)}})]})}),(0,a.jsx)("div",{className:"col-12 d-flex justify-content-center",children:(0,a.jsx)("button",{type:"submit",className:"btn btn-success btn-lg",children:"Submit"})}),(0,a.jsx)("div",{className:"col-12 d-flex justify-content-center",children:(0,a.jsx)(p(),{href:"/signup",children:(0,a.jsx)("a",{children:"No account? Sign up here!"})})}),y&&(0,a.jsx)("div",{className:"alert alert-danger",children:y})]})})]})]})}}},function(e){e.O(0,[452,664,774,888,179],(function(){return n=6429,e(e.s=n);var n}));var n=e.O();_N_E=n}]);