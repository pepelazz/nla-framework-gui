export default function (project) {
  let p = JSON.parse(JSON.stringify(project))
  p.docs = p.docs.map(doc => {
    doc.flds = doc.flds.map(fld => {
       // в параметрах image CanAddUrls из bool -> string
      if (fld.fld_vue_img_params) {
        fld.fld_vue_img_params.CanAddUrls = fld.fld_vue_img_params.CanAddUrls === true ? 'true' : 'false'
      }
      return fld
    })
    return doc
  })
  return p
}
