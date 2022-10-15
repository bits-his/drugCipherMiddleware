package db

const (
	cREATE_TABLE_DRUGS = `CREATE TABLE drugs(
		id int(11) NOT NULL,
		drug_name varchar(50) NOT NULL,
		drug_generic_name varchar(50) NOT NULL,
		dosage int(11) NOT NULL,
		date_enrolled timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
		expiry_date timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
		nafdac varchar(100) NOT NULL,
		qr_code string NOT NULL,
		manufacturers_id int(11) NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=latin1;`

	cREATE_TABLE_MANUFACTURERS = `CREATE TABLE manufacturers (
		id int(11) NOT NULL,
		pharmacy_license varchar(50) NOT NULL,
		pharmacist_license varchar(50) NOT NULL,
		name varchar(50) NOT NULL,
		address varchar(50) NOT NULL,
		po_box varchar(50) NOT NULL,
		email varchar(50) NOT NULL,
		phone_number varchar(50) NOT NULL
	  ) ENGINE=InnoDB DEFAULT CHARSET=latin1;`

	aLL_MANUFACTURERS    = `SELECT * FROM manufacturers;`
	aLL_DRUGS            = `SELECT * FROM drugs;`
	cREATE_MANUFACTURERS = `INSERT INTO manufacturers(
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
	cREATE_DRUGS = `
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
	gET_MANUFACTURER = `SELECT * FROM manufacturers WHERE id = ?`
	gET_DRUG         = `SELECT * FROM drugs WHERE id = ?`
	uPDATE_DRUG      = `
		UPDATE drugs
		set drug_name=_drug_name,
		drug_generic_name=_drug_generic_name,
		dosage=_dosage,
		date_enrolled=_date_enrolled,
		expiry_date=_expiry_date,
		nafdac=_nafdac,
		qr_code=_qrcode,
		Manufacturers_id=_manufacturers_id
			
		WHERE id=_id;
	`
	uPDATE_MANUFACTURER = `
	UPDATE manufacturers
	set pharmacy_license=_pharmacy_license,
	pharmacist_license=_pharmacist_license,
	name=_name,
	address=_address,
	po_box=_po_box,
	email=_email,
	phone_number=_phone_number

	WHERE id=_id;
	`
)
