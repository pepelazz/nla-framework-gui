import _ from "lodash";

export default (fld) => {
  switch (fld.func_name) {
    case 'GetFldTitle':
      return 'title'
    case 'GetFldRef':
      return 'link'
    case 'GetFldFiles':
      return 'file_present'
    case 'GetFldTag':
      return 'sell'
    case 'GetFldDate':
      return 'event'
    case 'GetFldDateTime':
      return 'schedule'
    case 'GetFldCheckbox':
      return 'check_circle_outline'
    case 'GetFldEmail':
      return 'email'
    case 'GetFldPhone':
      return 'phone'
    default:
      return 'text_fields'
  }
}
