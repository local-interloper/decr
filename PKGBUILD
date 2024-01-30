# Maintainer: local.interloper <marinoljubetic8@gmail.com>
pkgname=decr-git # '-bzr', '-git', '-hg' or '-svn'
pkgver=r4.3937898
pkgrel=1
pkgdesc="A tool for creating .desktop entries"
arch=('any')
url="https://github.com/local-interloper/decr"
license=('GPL')
groups=()
depends=()
makedepends=('git'
             'go') 
provides=("${pkgname%-git}")
conflicts=("${pkgname%-git}")
replaces=()
backup=()
options=()
install=
source=("git+${url}")
noextract=()
md5sums=('SKIP')

pkgver() {
	cd "$srcdir/${pkgname%-git}"
	printf "r%s.%s" "$(git rev-list --count HEAD)" "$(git rev-parse --short HEAD)"
}


build() {
	cd "$srcdir/${pkgname%-git}"
  go build -o decr
}

package() {
	cd "$srcdir/${pkgname%-git}"
  install -Dm755 decr "$pkgdir/usr/bin/decr"
}
