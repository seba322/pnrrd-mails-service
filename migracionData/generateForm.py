import json


#read json
f = open('form.json',)

data = json.load(f)

f.close()
fomrData = data["details"]["formData"]

newJson ={
    "_id": "adasda",
  "creation_date": 1552482488.87651,
  "modified_date": 1634831922.24348,
  "sections":[]
}

i=0
for key in fomrData:
    
    
    if i >1 and i<8:
        print(i,key)
        elForm = fomrData[key][0]["rows"]
        newObcj={
        "name":"Medios de Transporte",
        "hierarchy_type":"GENERAL",
        "resource_type":"Recurso Tecnico",
        "label":"Medios de Transporte",
        "index":i,
        "form":[]

            }
        j=0
        for  e in elForm:
            indeObj ={
                "label":e["label"],
                "index":str(i)+"-"+str(j)
            }
            
            newObcj["form"].append(indeObj)  
            j+=1 
        newJson["sections"].append(newObcj)
    i+=1

with open('forData.json', 'w', encoding='utf-8') as f:
    json.dump(newJson, f, ensure_ascii=False, indent=4)
#generate new json