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
    case 'GetFldInt64':
      return 'pin'
    case 'GetFldDouble':
      return 'pin'
    case 'GetFldSelectString':
      return 'rule'
    case 'GetFldSelectMultiple':
      return 'checklist'
    case 'GetFldRadioString':
      return 'radio_button_checked'
    case 'GetFldImg':
      return 'image'
    case 'GetFldImgList':
      return 'photo_library'
    default:
      return 'text_fields'
  }
}
