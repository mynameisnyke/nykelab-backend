gcloud functions deploy probeFile \
  --entry-point ProbeFile \
  --runtime go116 \
  --trigger-event "providers/cloud.firestore/eventTypes/document.create" \
  --trigger-resource "projects/nykelab/databases/(default)/documents/assets/{pushId}"