package firebase

type FirestoreAsset struct {
	ID          string   `firestore:"id" json:"id"`
	Type        string   `firestore:"type" json:"type"`
	Orientation string   `firestore:"orientation" json:"orientation"`
	Tags        []string `firestore:"tags,omitempty" json:"tags,omitempty"`
	Location    string   `firestore:"location" json:"location"`
	Dark        bool     `firestore:"displayName" json:"displayName"`
	Date        string   `firestore:"username" json:"username"`
}

type NewAsset struct {
	ID          string `firestore:"id" json:"id"`
	Name        string `firestore:"name" json:"name"`
	Type        string `firestore:"type" json:"type"`
	Upload_Date string `firestore:"upload_date" json:"upload_date"`
	Src         string `firestore:"src" json:"src"`
}

type VideoAssetProbeUpdate struct {
	ID            string `firestore:"id" json:"id"`
	Type          string `firestore:"type" json:"type"`
	Orientation   string `firestore:"orientation" json:"orientation"`
	Creation_Date string `firestore:"creation_date" json:"creation_date"`
}
