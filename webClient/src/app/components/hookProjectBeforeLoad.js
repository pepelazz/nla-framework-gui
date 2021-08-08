export default function (p) {
  p.docs = p.docs.map(doc => {
    doc.flds = doc.flds.map(fld => {
       // в параметрах image CanAddUrls из string -> bool
      if (fld.fld_vue_img_params) {
        fld.fld_vue_img_params.CanAddUrls = fld.fld_vue_img_params.CanAddUrls === 'true'
      }
      return fld
    })
    return doc
  })
  return p
}
