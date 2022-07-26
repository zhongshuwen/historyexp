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
const query : any= parseQuery(window.location.href)
const apiInURL = query.api_endpoint;
const secureInURL = query.api_endpoint_secure === "true"||query.api_endpoint_secure === "1";
const overrideConfig: any = apiInURL?{
  dfuse_io_endpoint: apiInURL,
  secure: secureInURL,
}:{}
const networkOverride: any = {
  is_test: false,
}

export {
  overrideConfig,
  networkOverride,
}