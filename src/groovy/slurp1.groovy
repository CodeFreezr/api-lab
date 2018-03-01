//import groovy.json.JsonSlurper
import java.util.zip.GZIPInputStream
def slurper = new groovy.json.JsonSlurper()
def a = "https://api.stackexchange.com/2.2/questions?filter=total&tagged=graphviz&site=stackoverflow".toURL().getText()
//def a = "https://httpbin.org/ip".toURL()

//print a.getClass()


def s = unzip(a)

//def s = slurper.parse(a)

//println s

//def url = "http://www.mrhaki.com/url.html".toURL()




def unzip(String compressed){
	def inflaterStream = new GZIPInputStream(new ByteArrayInputStream(compressed))
    def uncompressedStr = inflaterStream.getText('UTF-8')
    return uncompressedStr
}