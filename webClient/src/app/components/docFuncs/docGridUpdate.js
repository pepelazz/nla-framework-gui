export default (doc, flds) => {
  flds.map(v => {
    if (v.name === 'title') {
      doc.value.flds[0].col_class = v.col_class
    } else {
      let fld = doc.value.flds.find(v1 => v1.name === v.name)
      if (fld) {
        fld.col_class = v.col_class
        fld.row_col = v.row_col
      }
    }
  })
  // сортируем массив doc.flds в соответствии с новым порядком
  doc.value.flds.sort((a, b) => flds.findIndex(v => v.name === a.name) - flds.findIndex(v => v.name === b.name))
}
