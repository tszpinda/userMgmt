import DS from 'ember-data';

export default DS.Model.extend({
  name:     DS.attr('string'),
  domain:   DS.attr('string'),
  page: 	DS.attr('string'),
  apiKey:   DS.attr('string'),
  steps: DS.hasMany('step', { embedded: true, async: false })
});
