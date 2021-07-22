import _ from "lodash";

export default (flds) => _.chain(flds)
  // группируем по строкам
  .groupBy(v => v.row_col ? v.row_col[0][0] : 1)
  // внутри строки сортируем по колонкам
  .forEach(row => row.sort((a, b) => {
    return (a.row_col ? a.row_col[0][1] : 1) - (b.row_col ? b.row_col[0][1] : 1)
  }))
  .value()
