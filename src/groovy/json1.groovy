import groovy.json.JsonSlurper;

def test = '[{"cid":"PWER","data":[{"1440023704000":0}],"sid":"123456","units":"kWm","age":142156},{"cid":"PWER_GAC","data":[{"1440165858000":2202}],"sid":"123456","units":null,"age":2}]'
def myJson = /D:\dbt\01\awesome\awesome-graphviz\ab-results-README.md-filtered.json/
sFile = new File(myJson)
counter = 0

def respData = sFile.text
def list = new JsonSlurper().parseText(respData)
def cidval = "" // or whatever default you want
for (rec in list) {
    println counter++ rec.status
}