import DS from 'ember-data';

export default DS.Model.extend({
    no: 	DS.attr('number'),
    selector: 	DS.attr('string'),
  	text:   	DS.attr('string'),
  	tutorial: 	DS.belongsTo('tutorial', {embedded: false, async: false})
});
