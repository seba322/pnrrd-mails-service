from pymongo import MongoClient
import json
from bson import ObjectId, json_util
import datetime

client = MongoClient("localhost", 27017)


DB_USER = "pnrrd-prod"
DB_PASS = "clavepnrrd!"
DB_DB = "PNRRDPROD"
DB_URL = "localhost:27017"


def get_database():

    # Provide the mongodb atlas url to connect python to mongodb using pymongo
    # CONNECTION_STRING = "mongodb://<username>:<password>@<cluster-name>.mongodb.net/"

    CONNECTION_STRING = "mongodb://"+DB_USER+":"+DB_PASS+"@"+DB_URL+"/"+DB_DB
    # Create a connection using MongoClient. You can import MongoClient or use pymongo.MongoClient

    client = MongoClient(CONNECTION_STRING)

    # Create the database for our example (we will use the same database throughout the tutorial
    return client[DB_DB]


def readJsonData(name):
    f = open(name,)
    data = json.load(f)
    f.close()
    return data


def getAllDataCollection(collectionRef):
    # print("1")
    data = collectionRef.find({})
    # print("2")
    datArray = []
    # print("3")
    for d in data:
        # print(d)
        datArray.append(d)
    return datArray


def findDataRegionByid(json, id):

    for j in json:
        if j["id"] == id:
            return j["provincias"]


def addIdsJerarquias(data):

    for d in data:
        # gen_time = datetime.now()
        dummy_id = ObjectId()
        d["id"] = dummy_id
        for c in d["comunas"]:
            # gen_timeC = datetime.now()
            idC = ObjectId()
            c["id"] = idC


def reduceJearquias(jsonRegiones, regionesBd):
    nuevoJson = []
    for j in regionesBd:
        idD = j["_id"]
        if idD == 0:
            continue
        # print(idD)
        data = findDataRegionByid(jsonRegiones, idD)
        # print(data)
        addIdsJerarquias(data)
        j["provincias"] = data
        nuevoJson.append(j.copy())
    return nuevoJson


def insertArrayData(col, data):
    col.insert_many(data)


def getIndexForm(form, label, p):
    index = "-1"

    sections = form["sections"]

    mapOtro = {
        "paso4": "2-17",
        "paso5": "3-23",
        "paso6": "5-16",
        "paso7": "6-7",
        "paso8": "4-33",
        "paso9": "7-8",
        "paso3": "3G-0"
    }
    if label == "Otro":
        return mapOtro[p], ""

    defaultLabel = {
        "Centro de distribución (habilitación y administración)": "8-0",
        "Centro de acopio (habilitación y administración, bodegas por ejemplo)": "8-0"
    }
    newLabel = "Centro de acopio y distribución (habilitación y administración, bodegas)"
    if label in defaultLabel:
        return defaultLabel[label], newLabel
    for s in sections:
        for f in s["form"]:
            if "list" in f and label in f["list"]:
                index = f["index"]
                break

            if "options" in f and label in f["options"]:
                index = f["index"]
                break
            if f["label"] == label:
                index = f["index"]
                break

    return index, ""


def getNewInventory(inventoryData, generalForm, formData):

    GENERAL = ["paso1", "paso2", "paso3"]
    REGIONAL_HIERARCHY = "REGIONAL"
    NACIONAL_HIERARCHY = "NACIONAL"

    InformationTypeForm = "INFORMATION"

    ResourceTypeForm = "RESOURCE"

    rezagados = []

    newInventory = []
    nrezagados = 0
    nnew = 0
    labelRezagados = []

    for ins in inventoryData:
        insId = ins["institucion_id"]
        creationDate = ins["creation_date"]
        modifiedDate = ins["modified_date"]

        form = ins["details"]["formData"]

        for p in form:
            typeInv = ResourceTypeForm
            jerarquia = ""
            if p in GENERAL:
                typeInv = InformationTypeForm
                jerarquia = NACIONAL_HIERARCHY

            pasoArr = form[p]
            # print(p)
            if typeInv == InformationTypeForm:

                regionId = 0
                values = []
                if p == "paso1":
                    values = pasoArr.values()
                else:
                    values = pasoArr
                for c in values:
                    label = c["label"]

                    index = "-1"
                    if typeInv == InformationTypeForm:
                        index, _ = getIndexForm(generalForm, label, p)
                        if label == "" and c["value"] != "":
                            index = "1G-1"
                            c["label"] = "Misión organismo"
                    else:
                        index, _ = getIndexForm(formData, label, p)
                    newObjCap = {}
                    newObjCap["institucion_id"] = insId
                    newObjCap["creation_date"] = creationDate
                    newObjCap["modified_date"] = modifiedDate
                    newObjCap["type_inventory"] = typeInv
                    newObjCap["hierarchy"] = jerarquia
                    newObjCap["hierarchy_id"] = regionId
                    newObjCap["index"] = index
                    newObjCap["_id"] = ObjectId()
                    newObjCap["details"] = c
                    if index == "-1":
                        rezagados.append(newObjCap)
                        nrezagados += 1
                        if label not in labelRezagados:
                            labelRezagados.append(label)
                    else:
                        newInventory.append(newObjCap)
                        nnew += 1
            else:
                for r in pasoArr:

                    # print(r)

                    regionId = r["region"]

                    if regionId != 0:
                        jerarquia = REGIONAL_HIERARCHY
                    else:
                        jerarquia = NACIONAL_HIERARCHY

                    rowsArr = r["rows"]
                    if len(rowsArr) == 0:
                        continue
                    for cap in rowsArr:
                        if len(cap) <= 2:
                            continue

                        label = cap["label"]

                        index = "-1"
                        if typeInv == InformationTypeForm:
                            index, _ = getIndexForm(generalForm, label, p)
                        else:
                            index, newLabel = getIndexForm(formData, label, p)
                            label = newLabel

                        newObjCap = {}
                        newObjCap["institucion_id"] = insId
                        newObjCap["creation_date"] = creationDate
                        newObjCap["modified_date"] = modifiedDate
                        newObjCap["type_inventory"] = typeInv
                        newObjCap["hierarchy"] = jerarquia
                        newObjCap["hierarchy_id"] = regionId
                        newObjCap["index"] = index
                        newObjCap["_id"] = ObjectId()
                        newObjCap["details"] = cap

                        if index == "-1":
                            rezagados.append(newObjCap)
                            nrezagados += 1
                            if label not in labelRezagados:
                                labelRezagados.append(label)
                        else:
                            newInventory.append(newObjCap)
                            nnew += 1

    return newInventory, rezagados, nrezagados, nnew, labelRezagados


class JSONEncoder(json.JSONEncoder):
    def default(self, o):
        if isinstance(o, ObjectId):
            return str(o)
        return json.JSONEncoder.default(self, o)


def writeJson(name, data):
    with open(name, 'w', encoding='utf-8') as f:

        # json.encode(analytics, cls=JSONEncoder)
        #
        json.dump(data, f, ensure_ascii=False, indent=4)


    # print("conect")
DB = get_database()
# print("col")
# col = DB["regions_model"]

# jsonRegiones = readJsonData("jerarquias.json")
# print("jsonn")
# regionesBd = getAllDataCollection(col)
# print("col-get", regionesBd)

# nuevasJerarquias = reduceJearquias(jsonRegiones, regionesBd)


# print("nuevaj", nuevasJerarquias)
colJerarquias = DB["Hierarchy"]
# insertArrayData(colJerarquias, nuevasJerarquias)

jerarquiasJson = getAllDataCollection(colJerarquias)


formData = readJsonData("formData.json")
generalForm = readJsonData("generalForm.json")


colInventory = DB["inventory_model"]
currentInventory = getAllDataCollection(colInventory)


newInv, rezagados, nrezagados, nnew, labelRe = getNewInventory(
    currentInventory, generalForm, formData)


print("Rezagados:", nrezagados, " nnew:", nnew)
print("label:", labelRe)

colInventory = DB["NewInventory"]
colInventoryTest = DB["NewInventory_test"]
colInventoryTest2 = DB["NewInventory_test2"]

insertArrayData(colInventory, newInv)
insertArrayData(colInventoryTest, newInv)
insertArrayData(colInventoryTest2, newInv)
# writeJson("newinv.json", newInv)
# writeJson("rezagados.json", rezagados)
# This is added so that many files can reuse the function get_database()


# obtener json region

# leer arreglo regiones


# hacer match de regiones egun id y agregar provincias y comunas con sus respectivas id


# crear coleccion hierachy_model


# insertar data


# leer todas las instituciones


# por cada institucion crear un template


#
