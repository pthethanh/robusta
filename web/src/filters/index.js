import { formatDistance } from 'date-fns'
import es from 'date-fns/locale/es'
window.locale = es

export function date(v) {
  return formatDistance(new Date(v), new Date()) + ' ago'
}
