(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-4c57e8b0"],{"0798":function(t,e,i){"use strict";i("a57f");var n=i("9d26"),s=i("b64a"),r=i("98a1"),a=i("2b0e"),o=a["a"].extend({name:"transitionable",props:{mode:String,origin:String,transition:String}}),c=i("58df");e["a"]=Object(c["a"])(s["a"],r["a"],o).extend({name:"v-alert",props:{dismissible:Boolean,icon:String,outline:Boolean,type:{type:String,validator:function(t){return["info","error","success","warning"].includes(t)}}},computed:{computedColor:function(){return this.type&&!this.color?this.type:this.color||"error"},computedIcon:function(){if(this.icon||!this.type)return this.icon;switch(this.type){case"info":return"$vuetify.icons.info";case"error":return"$vuetify.icons.error";case"success":return"$vuetify.icons.success";case"warning":return"$vuetify.icons.warning"}}},methods:{genIcon:function(){return this.computedIcon?this.$createElement(n["a"],{class:"v-alert__icon"},this.computedIcon):null},genDismissible:function(){var t=this;return this.dismissible?this.$createElement("a",{class:"v-alert__dismissible",on:{click:function(){t.isActive=!1}}},[this.$createElement(n["a"],{props:{right:!0}},"$vuetify.icons.cancel")]):null}},render:function(t){var e=[this.genIcon(),t("div",this.$slots.default),this.genDismissible()],i=this.outline?this.setTextColor:this.setBackgroundColor,n=t("div",i(this.computedColor,{staticClass:"v-alert",class:{"v-alert--outline":this.outline},directives:[{name:"show",value:this.isActive}],on:this.$listeners}),e);return this.transition?t("transition",{props:{name:this.transition,origin:this.origin,mode:this.mode}},[n]):n}})},a57f:function(t,e,i){},b156:function(t,e,i){"use strict";i.r(e);var n=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("v-container",{attrs:{fluid:""}},[i("v-alert",{attrs:{value:null!=t.error,type:"error"}},[t._v(t._s(t.error))]),i("v-card",[i("v-data-table",{staticClass:"elevation-1",attrs:{headers:t.headers,items:t.items,"hide-actions":"true"},scopedSlots:t._u([{key:"items",fn:function(e){return[i("td",[t._v(t._s(e.item.Name))]),i("td",[t._v(t._s(e.item.Value))])]}}])})],1)],1)},s=[],r=i("9a1d"),a={name:"currentConfiguration",data:function(){return{headers:[{text:"Name",value:"Name"},{text:"Value",value:"Value"}],items:[]}},computed:{error:function(){return this.$store.state.error}},mounted:function(){var t=this;r["a"].getParameters().then(function(e){t.items=e.data},function(e){t.$store.dispatch("setError",e)})}},o=a,c=i("2877"),u=i("6544"),l=i.n(u),d=i("0798"),h=i("b0af"),f=i("a523"),m=i("8fea"),v=Object(c["a"])(o,n,s,!1,null,null,null);e["default"]=v.exports;l()(v,{VAlert:d["a"],VCard:h["a"],VContainer:f["a"],VDataTable:m["a"]})}}]);
//# sourceMappingURL=chunk-4c57e8b0.92acb813.js.map