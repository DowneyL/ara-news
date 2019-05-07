module ara

require (
	github.com/Unknwon/goconfig v0.0.0-20181105214110-56bd8ab18619 // indirect
	github.com/astaxie/beego v1.11.1
	github.com/beego/i18n v0.0.0-20161101132742-e9308947f407
	github.com/go-gomail/gomail v0.0.0-20160411212932-81ebce5c23df
	github.com/go-ini/ini v1.42.0
	github.com/go-playground/locales v0.12.1
	github.com/go-playground/universal-translator v0.16.0
	github.com/go-redis/redis v6.14.2+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/leodido/go-urn v1.1.0 // indirect
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/oschwald/geoip2-golang v1.2.1
	github.com/oschwald/maxminddb-golang v1.3.0 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190306220146-200a235640ff
	github.com/stretchr/testify v1.3.0 // indirect
	golang.org/x/sys v0.0.0-20190322080309-f49334f85ddc // indirect
	google.golang.org/appengine v1.4.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.27.0
	gopkg.in/ini.v1 v1.42.0 // indirect
)

replace (
	golang.org/x/net v0.0.0-20180724234803-3673e40ba225 => github.com/golang/net v0.0.0-20180724234803-3673e40ba225
	golang.org/x/net v0.0.0-20180906233101-161cd47e91fd => github.com/golang/net v0.0.0-20180906233101-161cd47e91fd
	golang.org/x/net v0.0.0-20181114220301-adae6a3d119a => github.com/golang/net v0.0.0-20181114220301-adae6a3d119a
)

replace golang.org/x/crypto v0.0.0-20181127143415-eb0de9b17e85 => github.com/golang/crypto v0.0.0-20181127143415-eb0de9b17e85

replace google.golang.org/appengine v1.4.0 => github.com/golang/appengine v1.4.0

replace golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0

replace (
	golang.org/x/sys v0.0.0-20180909124046-d0be0721c37e => github.com/golang/sys v0.0.0-20180909124046-d0be0721c37e
	golang.org/x/sys v0.0.0-20190322080309-f49334f85ddc => github.com/golang/sys v0.0.0-20190322080309-f49334f85ddc
)

replace golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f => github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f
