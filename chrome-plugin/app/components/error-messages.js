import Ember from 'ember';

/*
The ErrorMessagesComponent displays any validation errrors returned
from the API
 */
export default Ember.Component.extend({
	errors: Ember.computed.alias('for.errors')
});
