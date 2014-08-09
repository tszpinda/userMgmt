import DS from 'ember-data';

export default DS.Model.extend({
 	selector: 	DS.attr('string'),
  	text:   		DS.attr('string'),
  	tutorial: DS.belongsTo('tutorial')
});
