
use PNRRDPROD;

//crea usuario
db.dropUser("pnrrd-prod");

db.dropDatabase();

db.createUser({
	user: 'pnrrd-prod', pwd: 'clavepnrrd!',
	roles: [{ role: "readWrite", db: "PNRRDPROD" }]
});

//borra base de datos
db.getCollection('activation_model').drop();
db.getCollection('institucion_model').drop();
db.getCollection('inventory_model').drop();
db.getCollection('regions_model').drop();
db.getCollection('rol_model').drop();
db.getCollection('user_model').drop();
db.getCollection('NewInventory').drop();
db.getCollection('Hierachy').drop();
db.getCollection('Form').drop();

exit;
