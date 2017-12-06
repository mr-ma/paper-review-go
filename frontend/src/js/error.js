	function bootstrapShow( msg ) {
		BootstrapDialog.show({
          message: '' + msg
        });
	}

	function handleError ( error ) {
		if (!error) return;
		if (!error.code) {
			bootstrapShow(error);
		} else {
			switch (error.code) {
				case 'ER_CUSTOM_USER_NOT_FOUND': bootstrapShow('User email not found.'); break;
				case 'ER_CUSTOM_DUPLICATE_REGION': bootstrapShow('duplicate - not allowed - please rename the Region'); break;
				case 'ER_CUSTOM_DUPLICATE_USER': bootstrapShow('duplicate - not allowed - user already exists'); break;
				case 'ER_DUP_ENTRY': bootstrapShow('This entry already exists in the database. Please choose a different one.'); break;
				case 'ER_CUSTOM_NO_TITLE': bootstrapShow('Please enter a title.'); break;
				case 'ER_CUSTOM_NON_EXISTANT_ID': bootstrapShow('Wrong ID.'); break;
				case 'ER_CUSTOM_OUTLOOK_REJECT': bootstrapShow(error.msg); break;
				case 'ER_CUSTOM_DB_CONNECTION': bootstrapShow('Cannot connect to the database server.'); break;
				case 'PROTOCOL_ENQUEUE_AFTER_FATAL_ERROR': bootstrapShow(error.code); break;
				default: bootstrapShow(error.code);
			}
		}
	}