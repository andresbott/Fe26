import{s as k,Z as b,a as w,b as A,c as R,d as K,_ as O}from"./index-D7JBfKef.js";import{B as L,o as a,e as d,f as m,m as o,t as C,c as h,n as B,g as S,h as f,i as q,j as P,k as V,l as Z,F as T,r as g,p as j,w as c,a as l,T as N,q as U,s as $,b as v,u as D,v as F,x as s,_ as M,y as X}from"./index-CLsQzkYt.js";var Y=function(n){var t=n.dt;return`
.p-avatar {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: `.concat(t("avatar.width"),`;
    height: `).concat(t("avatar.height"),`;
    font-size: `).concat(t("avatar.font.size"),`;
    background: `).concat(t("avatar.background"),`;
    border-radius: `).concat(t("avatar.border.radius"),`;
}

.p-avatar-image {
    background: transparent;
}

.p-avatar-circle {
    border-radius: 50%;
}

.p-avatar-circle img {
    border-radius: 50%;
}

.p-avatar-icon {
    font-size: `).concat(t("avatar.font.size"),`;
}

.p-avatar img {
    width: 100%;
    height: 100%;
}

.p-avatar-lg {
    width: `).concat(t("avatar.lg.width"),`;
    height: `).concat(t("avatar.lg.width"),`;
    font-size: `).concat(t("avatar.lg.font.size"),`;
}

.p-avatar-lg .p-avatar-icon {
    font-size: `).concat(t("avatar.lg.font.size"),`;
}

.p-avatar-xl {
    width: `).concat(t("avatar.xl.width"),`;
    height: `).concat(t("avatar.xl.width"),`;
    font-size: `).concat(t("avatar.xl.font.size"),`;
}

.p-avatar-xl .p-avatar-icon {
    font-size: `).concat(t("avatar.xl.font.size"),`;
}

.p-avatar-group {
    display: flex;
    align-items: center;
}

.p-avatar-group .p-avatar + .p-avatar {
    margin-left: `).concat(t("avatar.group.offset"),`;
}

.p-avatar-group .p-avatar {
    border: 2px solid `).concat(t("avatar.group.border.color"),`;
}
`)},G={root:function(n){var t=n.props;return["p-avatar p-component",{"p-avatar-image":t.image!=null,"p-avatar-circle":t.shape==="circle","p-avatar-lg":t.size==="large","p-avatar-xl":t.size==="xlarge"}]},label:"p-avatar-label",icon:"p-avatar-icon"},H=L.extend({name:"avatar",theme:Y,classes:G}),J={name:"BaseAvatar",extends:k,props:{label:{type:String,default:null},icon:{type:String,default:null},image:{type:String,default:null},size:{type:String,default:"normal"},shape:{type:String,default:"square"},ariaLabelledby:{type:String,default:null},ariaLabel:{type:String,default:null}},style:H,provide:function(){return{$pcAvatar:this,$parentInstance:this}}},y={name:"Avatar",extends:J,inheritAttrs:!1,emits:["error"],methods:{onError:function(n){this.$emit("error",n)}}},Q=["aria-labelledby","aria-label"],W=["src","alt"];function _(e,n,t,u,i,r){return a(),d("div",o({class:e.cx("root"),"aria-labelledby":e.ariaLabelledby,"aria-label":e.ariaLabel},e.ptmi("root")),[m(e.$slots,"default",{},function(){return[e.label?(a(),d("span",o({key:0,class:e.cx("label")},e.ptm("label")),C(e.label),17)):e.$slots.icon?(a(),h(S(e.$slots.icon),{key:1,class:B(e.cx("icon"))},null,8,["class"])):e.icon?(a(),d("span",o({key:2,class:[e.cx("icon"),e.icon]},e.ptm("icon")),null,16)):e.image?(a(),d("img",o({key:3,src:e.image,alt:e.ariaLabel,onError:n[0]||(n[0]=function(){return r.onError&&r.onError.apply(r,arguments)})},e.ptm("image")),null,16,W)):f("",!0)]})],16,Q)}y.render=_;var ee=function(n){var t=n.dt;return`
.p-drawer {
    display: flex;
    flex-direction: column;
    pointer-events: auto;
    transform: translate3d(0px, 0px, 0px);
    position: relative;
    transition: transform 0.3s;
    background: `.concat(t("drawer.background"),`;
    color: `).concat(t("drawer.color"),`;
    border: 1px solid `).concat(t("drawer.border.color"),`;
    box-shadow: `).concat(t("drawer.shadow"),`;
}

.p-drawer-content {
    overflow-y: auto;
    flex-grow: 1;
    padding: `).concat(t("drawer.content.padding"),`;
}

.p-drawer-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    flex-shrink: 0;
    padding: `).concat(t("drawer.header.padding"),`;
}

.p-drawer-footer {
    padding: `).concat(t("drawer.header.padding"),`;
}

.p-drawer-title {
    font-weight: `).concat(t("drawer.title.font.weight"),`;
    font-size: `).concat(t("drawer.title.font.size"),`;
}

.p-drawer-full .p-drawer {
    transition: none;
    transform: none;
    width: 100vw !important;
    height: 100vh !important;
    max-height: 100%;
    top: 0px !important;
    left: 0px !important;
    border-width: 1px;
}

.p-drawer-left .p-drawer-enter-from,
.p-drawer-left .p-drawer-leave-to {
    transform: translateX(-100%);
}

.p-drawer-right .p-drawer-enter-from,
.p-drawer-right .p-drawer-leave-to {
    transform: translateX(100%);
}

.p-drawer-top .p-drawer-enter-from,
.p-drawer-top .p-drawer-leave-to {
    transform: translateY(-100%);
}

.p-drawer-bottom .p-drawer-enter-from,
.p-drawer-bottom .p-drawer-leave-to {
    transform: translateY(100%);
}

.p-drawer-full .p-drawer-enter-from,
.p-drawer-full .p-drawer-leave-to {
    opacity: 0;
}

.p-drawer-full .p-drawer-enter-active,
.p-drawer-full .p-drawer-leave-active {
    transition: opacity 400ms cubic-bezier(0.25, 0.8, 0.25, 1);
}

.p-drawer-left .p-drawer {
    width: 20rem;
    height: 100%;
    border-right-width: 1px;
}

.p-drawer-right .p-drawer {
    width: 20rem;
    height: 100%;
    border-left-width: 1px;
}

.p-drawer-top .p-drawer {
    height: 10rem;
    width: 100%;
    border-bottom-width: 1px;
}

.p-drawer-bottom .p-drawer {
    height: 10rem;
    width: 100%;
    border-top-width: 1px;
}

.p-drawer-left .p-drawer-content,
.p-drawer-right .p-drawer-content,
.p-drawer-top .p-drawer-content,
.p-drawer-bottom .p-drawer-content {
    width: 100%;
    height: 100%;
}

.p-drawer-open {
    display: flex;
}
`)},te={mask:function(n){var t=n.position;return{position:"fixed",height:"100%",width:"100%",left:0,top:0,display:"flex",justifyContent:t==="left"?"flex-start":t==="right"?"flex-end":"center",alignItems:t==="top"?"flex-start":t==="bottom"?"flex-end":"center"}}},ne={mask:function(n){var t=n.instance,u=n.props,i=["left","right","top","bottom"],r=i.find(function(p){return p===u.position});return["p-drawer-mask",{"p-overlay-mask p-overlay-mask-enter":u.modal,"p-drawer-open":t.containerVisible,"p-drawer-full":t.fullScreen},r?"p-drawer-".concat(r):""]},root:function(n){var t=n.instance;return["p-drawer p-component",{"p-drawer-full":t.fullScreen}]},header:"p-drawer-header",title:"p-drawer-title",pcCloseButton:"p-drawer-close-button",content:"p-drawer-content",footer:"p-drawer-footer"},re=L.extend({name:"drawer",theme:ee,classes:ne,inlineStyles:te}),ae={name:"BaseDrawer",extends:k,props:{visible:{type:Boolean,default:!1},position:{type:String,default:"left"},header:{type:null,default:null},baseZIndex:{type:Number,default:0},autoZIndex:{type:Boolean,default:!0},dismissable:{type:Boolean,default:!0},showCloseIcon:{type:Boolean,default:!0},closeButtonProps:{type:Object,default:function(){return{severity:"secondary",text:!0,rounded:!0}}},closeIcon:{type:String,default:void 0},modal:{type:Boolean,default:!0},blockScroll:{type:Boolean,default:!1}},style:re,provide:function(){return{$pcDrawer:this,$parentInstance:this}},watch:{dismissable:function(n){n?this.bindOutsideClickListener():this.unbindOutsideClickListener()}}},x={name:"Drawer",extends:ae,inheritAttrs:!1,emits:["update:visible","show","hide","after-hide"],data:function(){return{containerVisible:this.visible}},container:null,mask:null,content:null,headerContainer:null,footerContainer:null,closeButton:null,outsideClickListener:null,documentKeydownListener:null,updated:function(){this.visible&&(this.containerVisible=this.visible)},beforeUnmount:function(){this.disableDocumentSettings(),this.mask&&this.autoZIndex&&b.clear(this.mask),this.container=null,this.mask=null},methods:{hide:function(){this.$emit("update:visible",!1)},onEnter:function(){this.$emit("show"),this.focus(),this.bindDocumentKeyDownListener(),this.autoZIndex&&b.set("modal",this.mask,this.baseZIndex||this.$primevue.config.zIndex.modal)},onAfterEnter:function(){this.enableDocumentSettings()},onBeforeLeave:function(){this.modal&&!this.isUnstyled&&q(this.mask,"p-overlay-mask-leave")},onLeave:function(){this.$emit("hide")},onAfterLeave:function(){this.autoZIndex&&b.clear(this.mask),this.unbindDocumentKeyDownListener(),this.containerVisible=!1,this.disableDocumentSettings(),this.$emit("after-hide")},onMaskClick:function(n){this.dismissable&&this.modal&&this.mask===n.target&&this.hide()},focus:function(){var n=function(i){return i&&i.querySelector("[autofocus]")},t=this.$slots.header&&n(this.headerContainer);t||(t=this.$slots.default&&n(this.container),t||(t=this.$slots.footer&&n(this.footerContainer),t||(t=this.closeButton))),t&&P(t)},enableDocumentSettings:function(){this.dismissable&&!this.modal&&this.bindOutsideClickListener(),this.blockScroll&&V()},disableDocumentSettings:function(){this.unbindOutsideClickListener(),this.blockScroll&&Z()},onKeydown:function(n){n.code==="Escape"&&this.hide()},containerRef:function(n){this.container=n},maskRef:function(n){this.mask=n},contentRef:function(n){this.content=n},headerContainerRef:function(n){this.headerContainer=n},footerContainerRef:function(n){this.footerContainer=n},closeButtonRef:function(n){this.closeButton=n?n.$el:void 0},bindDocumentKeyDownListener:function(){this.documentKeydownListener||(this.documentKeydownListener=this.onKeydown,document.addEventListener("keydown",this.documentKeydownListener))},unbindDocumentKeyDownListener:function(){this.documentKeydownListener&&(document.removeEventListener("keydown",this.documentKeydownListener),this.documentKeydownListener=null)},bindOutsideClickListener:function(){var n=this;this.outsideClickListener||(this.outsideClickListener=function(t){n.isOutsideClicked(t)&&n.hide()},document.addEventListener("click",this.outsideClickListener))},unbindOutsideClickListener:function(){this.outsideClickListener&&(document.removeEventListener("click",this.outsideClickListener),this.outsideClickListener=null)},isOutsideClicked:function(n){return this.container&&!this.container.contains(n.target)}},computed:{fullScreen:function(){return this.position==="full"},closeAriaLabel:function(){return this.$primevue.config.locale.aria?this.$primevue.config.locale.aria.close:void 0}},directives:{focustrap:T},components:{Button:w,Portal:A,TimesIcon:R}},oe=["aria-modal"];function ie(e,n,t,u,i,r){var p=g("Button"),z=g("Portal"),E=j("focustrap");return a(),h(z,null,{default:c(function(){return[i.containerVisible?(a(),d("div",o({key:0,ref:r.maskRef,onMousedown:n[0]||(n[0]=function(){return r.onMaskClick&&r.onMaskClick.apply(r,arguments)}),class:e.cx("mask"),style:e.sx("mask",!0,{position:e.position})},e.ptm("mask")),[l(N,o({name:"p-drawer",onEnter:r.onEnter,onAfterEnter:r.onAfterEnter,onBeforeLeave:r.onBeforeLeave,onLeave:r.onLeave,onAfterLeave:r.onAfterLeave,appear:""},e.ptm("transition")),{default:c(function(){return[e.visible?U((a(),d("div",o({key:0,ref:r.containerRef,class:e.cx("root"),role:"complementary","aria-modal":e.modal},e.ptmi("root")),[e.$slots.container?m(e.$slots,"container",{key:0,closeCallback:r.hide}):(a(),d($,{key:1},[v("div",o({ref:r.headerContainerRef,class:e.cx("header")},e.ptm("header")),[m(e.$slots,"header",{class:B(e.cx("title"))},function(){return[e.header?(a(),d("div",o({key:0,class:e.cx("title")},e.ptm("title")),C(e.header),17)):f("",!0)]}),e.showCloseIcon?(a(),h(p,o({key:0,ref:r.closeButtonRef,type:"button",class:e.cx("pcCloseButton"),"aria-label":r.closeAriaLabel,unstyled:e.unstyled,onClick:r.hide},e.closeButtonProps,{pt:e.ptm("pcCloseButton"),"data-pc-group-section":"iconcontainer"}),{icon:c(function(I){return[m(e.$slots,"closeicon",{},function(){return[(a(),h(S(e.closeIcon?"span":"TimesIcon"),o({class:[e.closeIcon,I.class]},e.ptm("pcCloseButton").icon),null,16,["class"]))]})]}),_:3},16,["class","aria-label","unstyled","onClick","pt"])):f("",!0)],16),v("div",o({ref:r.contentRef,class:e.cx("content")},e.ptm("content")),[m(e.$slots,"default")],16),v("div",o({ref:r.footerContainerRef,class:e.cx("footer")},e.ptm("footer")),[m(e.$slots,"footer")],16)],64))],16,oe)),[[E]]):f("",!0)]}),_:3},16,["onEnter","onAfterEnter","onBeforeLeave","onLeave","onAfterLeave"])],16)):f("",!0)]}),_:3})}x.render=ie;const se=v("p",null,"TODO make to look more like github",-1),le={__name:"UserProfile",setup(e){const n=D(),t=F(!1),u=()=>{n.logout()};return(i,r)=>(a(),d($,null,[l(s(y),{icon:"pi pi-user",class:"mr-2",size:"large",onClick:r[0]||(r[0]=p=>t.value=!0),style:{"background-color":"#ece9fc",color:"#2a1261",cursor:"pointer"}}),l(s(x),{visible:t.value,"onUpdate:visible":r[3]||(r[3]=p=>t.value=p),header:s(n).userName,style:{width:"25rem"},position:"right"},{default:c(()=>[l(s(y),{icon:"pi pi-user",class:"mr-3",size:"xlarge",onClick:r[1]||(r[1]=p=>t.value=!0),style:{"background-color":"#ece9fc",color:"#2a1261"}}),se,v("p",null,[l(s(w),{label:"Settings",icon:"pi pi-cog"})]),v("p",null,[l(s(w),{label:"Logout",severity:"danger",icon:"pi pi-sign-out",onClick:r[2]||(r[2]=p=>u())})])]),_:1},8,["visible","header"])],64))}},de={},ce={width:"40",height:"40",viewBox:"0 0 50 50"},ue=X('<g id="layer1"><g id="g2148" transform="matrix(0.35823795,0,0,0.35823795,-10.823795,-10.823795)"><rect style="fill:#0090bd;fill-opacity:1;stroke:#2f3b3d;stroke-width:0;stroke-dasharray:none;stroke-opacity:1;" id="rect111" width="135.57204" height="135.57204" x="32.213982" y="32.213982" ry="10.678464"></rect><g aria-label="C" id="text1510" style="fill:#2f3b3d;"><path d="m 103.88005,73.120513 q -8.846129,0 -13.698864,6.67251 -4.852734,6.62196 -4.852734,18.50105 0,24.718617 18.551598,24.718617 7.7846,0 18.8549,-3.8923 v 13.14282 q -9.09888,3.7912 -20.32083,3.7912 -16.125231,0 -24.668066,-9.75602 -8.542834,-9.80656 -8.542834,-28.105416 0,-11.525244 4.195593,-20.169177 4.195593,-8.694483 12.030737,-13.294471 7.885694,-4.650537 18.4505,-4.650537 10.76701,0 21.63511,5.20658 l -5.05493,12.738428 q -4.14505,-1.971424 -8.34064,-3.437354 -4.19559,-1.46593 -8.23954,-1.46593 z" id="path1573"></path></g><g aria-label="5" id="text1668" style="fill:#2f3b3d;stroke-width:0;"><path d="m 145.7815,51.835121 q 4.27556,0 6.79654,2.399965 2.54114,2.399965 2.54114,6.574694 0,4.941105 -3.04534,7.603251 -3.04533,2.662146 -8.71248,2.662146 -4.92093,0 -7.9461,-1.593254 v -5.384796 q 1.59325,0.847047 3.71087,1.391577 2.11762,0.524362 4.01339,0.524362 5.70748,0 5.70748,-4.678924 0,-4.457078 -5.90916,-4.457078 -1.06889,0 -2.35963,0.221846 -1.29074,0.201678 -2.09745,0.443691 l -2.48063,-1.331073 1.10922,-15.024991 h 15.99305 v 5.283956 h -10.52758 l -0.54453,5.788151 0.70587,-0.141174 q 1.23024,-0.282349 3.04534,-0.282349 z" id="path1672"></path></g></g></g>',1),pe=[ue];function fe(e,n){return a(),d("svg",ce,pe)}const he=M(de,[["render",fe]]),we={__name:"topbar",setup(e){const n=D();return(t,u)=>{const i=g("router-link");return a(),h(O,{"center-content":!0},{left:c(()=>[l(i,{to:"/app",class:"layout-topbar-logo"},{default:c(()=>[l(he)]),_:1})]),right:c(()=>[s(n).isLoggedIn?(a(),h(le,{key:0})):f("",!0),s(n).isLoggedIn?f("",!0):(a(),h(i,{key:1,to:"/login",class:"layout-topbar-logo"},{default:c(()=>[l(s(w),{label:"Login",icon:"pi pi-sign-in"})]),_:1}))]),default:c(()=>[l(s(K),{placeholder:"Search",type:"text",class:"w-32 sm:w-auto"})]),_:1})}}};export{he as P,we as _};
