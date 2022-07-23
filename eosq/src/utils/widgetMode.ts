function parseQuery(url: string){
  const qInd = url.indexOf("?")
  if(qInd === -1){
    return {};
  }
  const qPart = url.substring(qInd);
  const outObj = {};
  qPart.split("&").filter(x=>x).forEach(x=>{
    const eqInd = x.indexOf("=")
    if (eqInd!==-1){
      outObj[decodeURIComponent(x.substring(0, eqInd))]= decodeURIComponent(x.substring(eqInd+1));
    }
  })
  return outObj;
}

const QueryMap : any = parseQuery(window.location.href);
const WIDGET_MODE : boolean = QueryMap.widget_mode === "1";
const WidgetTitle : string = QueryMap.widget_title || "";

function isWidgetModeActivated(): boolean {
  return WIDGET_MODE;
}
function getWidgetTitle() {
  return WidgetTitle;
}
function setupStyle() {
  if(isWidgetModeActivated()){
    document.body.style.background ="transparent !important";
    document.body.style.background ="transparent";
    const widgetOuter: any = document.getElementsByClassName("widgetOuter");
    if(widgetOuter && widgetOuter[0] && widgetOuter[0].children && widgetOuter[0].children[0]){
      const widgetModeChild : any= widgetOuter[0].children[0];
      widgetModeChild.style.padding = "0px"
      widgetModeChild.style.margin = "0px"
    }

    document.body.style.background ="transparent !important";
    document.body.style.background ="transparent";
    if(typeof window.postMessage==='function' && typeof window.parent.postMessage === 'function'){
      window.parent.postMessage(JSON.stringify({iframe_key: QueryMap.iframe_key, height:document.body.scrollHeight}), "*");
    }
  }
}
function runWidgetMode(){
  if(isWidgetModeActivated()){
    setInterval(()=>{
      setupStyle();
    },3000);
  }
}
runWidgetMode();
export {
  isWidgetModeActivated,
  getWidgetTitle
}