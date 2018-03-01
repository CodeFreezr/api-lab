def counter = 0
def start = System.currentTimeMillis()


//def http = new HTTPBuilder('http://myServer/')
//http.contentEncoding = ContentEncoding.Type.GZIP


def http = new HTTPBuilder('http://www.google.com')
def html = http.get( path : '/search', query : [q:'Groovy'] )

print html.getCLass()


//def url = new URL('https://api.stackexchange.com/2.2/questions/no-answers?order=desc&sort=activity&tagged=graphviz&site=stackoverflow')

//println url.getText()

/*
def list = new JsonSlurper().parseText(text)
for (rec in list) {
        println counter++ + ". " + rec.status + ": " + rec.link + " ->|-> " + rec.redirect
}


new File("") << urlInfo.getText()
*/