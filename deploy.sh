gcloud functions deploy probeFile \
  --entry-point ProbeFile \
  --runtime go116 \
  --trigger-event "providers/cloud.firestore/eventTypes/document.create" \
  --retry  \
  --trigger-resource "projects/nykelab/databases/(default)/documents/assets/{pushId}"

  gcloud functions deploy convertFile \
  --entry-point ConvertFile \
  --runtime go116 \
  --trigger-event "providers/cloud.firestore/eventTypes/document.create" \
  --retry  \
  --trigger-resource "projects/nykelab/databases/(default)/documents/assets/{pushId}"