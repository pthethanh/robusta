export function getTagType(tag) {
  let types = ['', 'info', 'success', 'warning', 'danger']
  let i = tag.length % types.length
  return types[i]
}

export function getTags(tags) {
  let result = []
  if (tags !== undefined) {
    if (Array.isArray(tags)) {
      for (var i = 0; i < tags.length; i++) {
        result.push({
          label: tags[i],
          type: getTagType(tags[i])
        })
      }
      return result
    }
    // it just a string
    result.push({
      label: tags,
      type: getTagType(tags)
    })
  }
  return result
}
