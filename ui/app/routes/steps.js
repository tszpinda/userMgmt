import Ember from 'ember';
import AuthRoute from './auth';

export default AuthRoute.extend({
  model: function() {
    /*see the hacks file for full explanation */
    //return hasManyHack(this.modelFo
console.log(this.modelFor('tutorial').get('name'));
    return this.modelFor('tutorial').get('steps');
    //return this.hasManyHack(this.modelFor('tutorial'), 'steps');
  },

  hasManyHack : function(parentRecord, collectionName) {
	  var parentId    = parentRecord.get('id');
	  var relationships   = Ember.get(parentRecord.constructor, 'relationshipsByName');
	  var relationship    = relationships.get(collectionName);
	  var idPath          = parentRecord.constructor.typeKey + '.id';

	  return parentRecord.get(collectionName).then(function() {
	    return parentRecord.store.filter(relationship.type, function(record) {
	      return record.get(idPath) === parentId;
	    });
	  });
	}

});
