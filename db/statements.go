package db

const (
	CREATE_TABLE_DRUGS = `CREATE TABLE drugs(
		id int(11) NOT NULL,
		drug_name varchar(50) NOT NULL,
		drug_generic_name varchar(50) NOT NULL,
		dosage int(11) NOT NULL,
		date_enrolled timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
		expiry_date timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
		nafdac varchar(100) NOT NULL,
		qr_code blob NOT NULL,
		manufacturers_id int(11) NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=latin1;`

	CREATE_TABLE_MANUFACTURERS = `CREATE TABLE manufacturers (
		manufacturers_id int(11) NOT NULL,
		pharmacy_license varchar(50) NOT NULL,
		pharmacist_license varchar(50) NOT NULL,
		name varchar(50) NOT NULL,
		address varchar(50) NOT NULL,
		po_box varchar(50) NOT NULL,
		email varchar(50) NOT NULL,
		phone_number varchar(50) NOT NULL
	  ) ENGINE=InnoDB DEFAULT CHARSET=latin1;`

	ALL_MANUFACTURERS    = `SELECT * FROM manufacturers;`
	ALL_DRUGS            = `SELECT * FROM drugs;`
	CREATE_MANUFACTURERS = `INSERT INTO manufacturers(
    	pharmacy_license,
        pharmacist_license,
        name,
        address,
        po_box,
        email,
        phone_number
    ) VALUES(
    	_pharmacy_license,
        _pharmacist_license,
        _name,
        _address,
        _po_box,
        _email,
        _phone_number
    );`
	CREATE_DRUGS = `
	INSERT INTO drugs (
    	drug_name,
        drug_generic_name,
        dosage,
        date_enrolled,
        expiry_date,
        nafdac,
        qr_code,
        manufacturers_id
        
    ) VALUES(
    	_drug_name,
        _drug_generic_name,
        _dosage,
        _date_enrolled,
        _expiry_date,
        _nafdac,
        _qr_code,
        _manufacturers_id
    );
	`
	GET_MANUFACTURER = `SELECT * FROM manufacturers WHERE manufacturer_id = ?`
	GET_DRUG         = `SELECT * FROM drugs WHERE id = ?`
)
