from pymongo import MongoClient
import json
from bson import ObjectId
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
        print(d)
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


    # print("conect")
DB = get_database()
# print("col")
col = DB["regions_model"]

jsonRegiones = readJsonData("jerarquias.json")
# print("jsonn")
regionesBd = getAllDataCollection(col)
# print("col-get", regionesBd)

nuevasJerarquias = reduceJearquias(jsonRegiones, regionesBd)


# print("nuevaj", nuevasJerarquias)
colJerarquias = DB["Hierarchy"]
insertArrayData(colJerarquias, nuevasJerarquias)
# This is added so that many files can reuse the function get_database()


# obtener json region

# leer arreglo regiones


# hacer match de regiones egun id y agregar provincias y comunas con sus respectivas id


# crear coleccion hierachy_model


# insertar data


# leer todas las instituciones


# por cada institucion crear un template


#
