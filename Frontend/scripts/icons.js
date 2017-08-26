const fs = require('fs')
const glob = require('glob')
var cheerio = require('cheerio')
var assign = require('object-assign')

function parse(contents, opt) {
  opt = opt || {}
  var $ = cheerio.load(contents, assign({ xmlMode: true }, opt))

  var fullpath = ''

  $('path').each(function() {
    var d = $(this).attr('d')
    fullpath += d.replace(/\s+/g, ' ')+' '
  })

  return {
    viewBox: $('svg').attr('viewBox'),
    paths: fullpath.trim()
  }
}

function extract(file, opt) {
  opt = opt || {}

  if (!opt.encoding) {
    opt.encoding = 'utf8'
  }

  var contents = fs.readFileSync(file, opt.encoding)
  return parse(contents, opt)
}

function getName(path) {
  const splited = path.split('/')
  return splited[splited.length-1].replace('.svg', '')
}

function gen(path, icons) {
  return glob.sync(`${path}/*.svg`).reduce((memo, current) => {
    const name = getName(current)
    const svg = extract(current)

    return Object.assign(memo, {[name]: svg})
  }, icons)
}

const icons = gen(__dirname + '/../assets/icons', {})

fs.writeFileSync(__dirname + '/../src/components/Icons/map.json', JSON.stringify(icons))
