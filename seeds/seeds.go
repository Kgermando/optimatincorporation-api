package seeds

import (
	"log"
	"time"

	"github.com/kgermando/optimatincorporation-api/models"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	seedUsers(db)
	seedServices(db)
	seedTeam(db)
	seedPortfolio(db)
	seedBlog(db)
}

// ── Users ─────────────────────────────────────────────────────────────────────

func seedUsers(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Count(&count)
	if count > 0 {
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte("Admin@optimat2025"), bcrypt.DefaultCost)
	if err != nil {
		log.Println("[seed] error hashing admin password:", err)
		return
	}

	users := []models.User{
		{Name: "Admin OPTIMAT", Email: "admin@optimat.cd", Password: string(hash), Role: "admin"},
	}

	if err := db.Create(&users).Error; err != nil {
		log.Println("[seed] error seeding users:", err)
		return
	}
	log.Println("[seed] users seeded")
}

// ── Services ──────────────────────────────────────────────────────────────────

func seedServices(db *gorm.DB) {
	var count int64
	db.Model(&models.Service{}).Count(&count)
	if count > 0 {
		return
	}

	services := []models.Service{
		{
			Title:       "Long-métrage",
			Subtitle:    "10 à 90 minutes",
			Description: "Films d'animation 3D captivants pour les enfants et familles, allant de 10 à 90 minutes.",
			Icon:        "🎬",
			Features:    pq.StringArray{"Scénario complet", "Animation 3D haute résolution", "Post-production", "Musique originale"},
			Order:       1,
			Published:   true,
		},
		{
			Title:       "Court-métrage",
			Subtitle:    "Illustration de projets",
			Description: "Nos courts-métrages vidéo illustrent parfaitement vos projets avec créativité et impact.",
			Icon:        "📺",
			Features:    pq.StringArray{"2 à 10 minutes", "Rendu photoréaliste", "Livraison rapide", "Tous formats"},
			Order:       2,
			Published:   true,
		},
		{
			Title:       "Moyen-métrage",
			Subtitle:    "Documentaires & simulations",
			Description: "Productions de moyen-métrages pour documentaires institutionnels et simulations 3D immersives.",
			Icon:        "🎥",
			Features:    pq.StringArray{"10 à 40 minutes", "Simulation réaliste", "Narration professionnelle", "Diffusion TV"},
			Order:       3,
			Published:   true,
		},
		{
			Title:       "Animation 3D",
			Subtitle:    "Création pure",
			Description: "Animations 3D innovantes pour donner vie à vos idées et projets avec un réalisme saisissant.",
			Icon:        "✨",
			Features:    pq.StringArray{"Personnages 3D", "Environnements", "Effets spéciaux", "VFX"},
			Order:       4,
			Published:   true,
		},
		{
			Title:       "Motion Graphique",
			Subtitle:    "Dynamisme visuel",
			Description: "Création de motion graphics pour des visuels dynamiques, attrayants et mémorables.",
			Icon:        "📊",
			Features:    pq.StringArray{"Infographies animées", "Titrage", "Transition", "Logo animation"},
			Order:       5,
			Published:   true,
		},
		{
			Title:       "Spots Publicitaires",
			Subtitle:    "Marketing impactant",
			Description: "Production de spots publicitaires percutants pour promouvoir efficacement vos produits.",
			Icon:        "📡",
			Features:    pq.StringArray{"15 à 60 secondes", "Concept créatif", "Musique et voix", "Multi-plateformes"},
			Order:       6,
			Published:   true,
		},
		{
			Title:       "Communication Institutionnelle",
			Subtitle:    "Image de marque",
			Description: "Supports de communication institutionnelle pour renforcer votre image et votre identité.",
			Icon:        "🏢",
			Features:    pq.StringArray{"Présentation entreprise", "Rapport annuel", "Brand film", "Visite virtuelle"},
			Order:       7,
			Published:   true,
		},
		{
			Title:       "Présentation de Projets",
			Subtitle:    "Convaincre et séduire",
			Description: "Supports visuels percutants pour des présentations de projets qui laissent une impression durable.",
			Icon:        "📋",
			Features:    pq.StringArray{"Visualisation architecturale", "Simulation technique", "Rendu 3D", "VR/AR"},
			Order:       8,
			Published:   true,
		},
		{
			Title:       "Documentaires Institutionnels",
			Subtitle:    "Racontez votre histoire",
			Description: "Production de documentaires pour raconter l'histoire et les valeurs de votre organisation.",
			Icon:        "🎞️",
			Features:    pq.StringArray{"Interviews", "Animation 3D", "Archives", "Diffusion"},
			Order:       9,
			Published:   true,
		},
	}

	if err := db.Create(&services).Error; err != nil {
		log.Println("[seed] error seeding services:", err)
		return
	}
	log.Println("[seed] services seeded")
}

// ── Team ──────────────────────────────────────────────────────────────────────

func seedTeam(db *gorm.DB) {
	var count int64
	db.Model(&models.TeamMember{}).Count(&count)
	if count > 0 {
		return
	}

	members := []models.TeamMember{
		{
			Name:     "Julio Fernando",
			Position: "Directeur Général",
			Bio:      "Fondateur et visionnaire d'OPTIMAT Incorporation, Julio Fernando dirige l'entreprise avec passion depuis 2016.",
			PhotoURL: "/images/teams/PHOTO OPTIMAT_Fernando.png",
			Order:    1,
		},
		{
			Name:     "Laeticia SHOMA",
			Position: "Directrice Administrative & Financière",
			Bio:      "Responsable de la gestion administrative et financière de l'entreprise.",
			PhotoURL: "/images/teams/PHOTO OPTIMAT_Laeticia.png",
			Order:    2,
		},
		{
			Name:     "Yannick Lufuluabo LFK",
			Position: "Directeur Production 3D",
			Bio:      "Responsable de toutes les productions 3D et de l'excellence technique des projets.",
			PhotoURL: "/images/teams/PHOTO OPTIMAT_Yannick.png",
			Order:    3,
		},
		{
			Name:     "Betty Stéphanie MBUYI",
			Position: "Directrice Études & Projets",
			Bio:      "En charge des études de projets et de la planification des productions.",
			PhotoURL: "/images/teams/PHOTO OPTIMAT_Betty_2.png",
			Order:    4,
		},
		{
			Name:     "Mbinza Odine",
			Position: "Directrice Artistique & Design",
			Bio:      "Créatrice de l'identité visuelle et artistique de toutes les productions OPTIMAT.",
			PhotoURL: "/images/teams/PHOTO OPTIMAT_Odine.png",
			Order:    5,
		},
		{
			Name:     "Loïc Imidy",
			Position: "Directeur Partenariats Commerciaux",
			Bio:      "Développe les partenariats stratégiques et les relations commerciales.",
			PhotoURL: "/images/teams/PHOTO OPTIMAT_Loic.png",
			Order:    6,
		},
		{
			Name:     "Mariana Dioma",
			Position: "Directrice de la Création",
			Bio:      "Supervise l'ensemble du processus créatif et de la direction artistique.",
			PhotoURL: "/images/teams/PHOTO OPTIMAT_Mariana Dioma.png",
			Order:    7,
		},
		{
			Name:     "Grâce MUBANGI",
			Position: "Directrice Juridique & Marchés",
			Bio:      "Gère les aspects juridiques, les contrats et la passation des marchés.",
			PhotoURL: "/images/teams/PHOTO OPTIMAT_Grace.png",
			Order:    8,
		},
	}

	if err := db.Create(&members).Error; err != nil {
		log.Println("[seed] error seeding team:", err)
		return
	}
	log.Println("[seed] team seeded")
}

// ── Portfolio ─────────────────────────────────────────────────────────────────

func seedPortfolio(db *gorm.DB) {
	var count int64
	db.Model(&models.Portfolio{}).Count(&count)
	if count > 0 {
		return
	}

	items := []models.Portfolio{
		{
			Title:       "Koffi Raymet — Saison 1",
			Slug:        "koffi-raymet-saison-1",
			Description: "La première série animée 3D congolaise, une aventure captivante pour les enfants et les familles.",
			Category:    "film",
			Year:        2022,
			Client:      "OPTIMAT Production",
			CoverURL:    "/images/3d/Koffi_16.jpg",
			GalleryURLs: pq.StringArray{},
			Tags:        pq.StringArray{"Film", "Série", "Enfants", "Animation 3D"},
			Published:   true,
		},
		{
			Title:       "Affiche Publicitaire 3D",
			Slug:        "affiche-publicitaire-3d",
			Description: "Création d'une affiche publicitaire en 3D pour une campagne marketing impactante.",
			Category:    "spot_publicitaire",
			Year:        2023,
			Client:      "Client Institutionnel",
			CoverURL:    "/images/3d/Affiche.jpg",
			GalleryURLs: pq.StringArray{},
			Tags:        pq.StringArray{"Publicité", "3D", "Affiche"},
			Published:   true,
		},
		{
			Title:       "Motion Graphics — Corporate",
			Slug:        "motion-graphics-corporate",
			Description: "Production de motion graphics pour une présentation corporate dynamique et mémorable.",
			Category:    "motion_graphics",
			Year:        2023,
			Client:      "Entreprise Locale",
			CoverURL:    "",
			GalleryURLs: pq.StringArray{},
			Tags:        pq.StringArray{"Motion", "Corporate", "Animation"},
			Published:   true,
		},
		{
			Title:       "Court-Métrage Animation",
			Slug:        "court-metrage-animation",
			Description: "Un court-métrage d'animation illustrant une histoire locale avec des personnages 3D expressifs.",
			Category:    "court_metrage",
			Year:        2022,
			Client:      "Producteur Indépendant",
			CoverURL:    "",
			GalleryURLs: pq.StringArray{},
			Tags:        pq.StringArray{"Film", "Animation", "Court-métrage"},
			Published:   true,
		},
		{
			Title:       "Documentaire Institutionnel",
			Slug:        "documentaire-institutionnel",
			Description: "Documentaire institutionnel retraçant l'histoire et les réalisations d'une organisation.",
			Category:    "documentaire",
			Year:        2023,
			Client:      "Organisation Institutionnelle",
			CoverURL:    "",
			GalleryURLs: pq.StringArray{},
			Tags:        pq.StringArray{"Documentaire", "Institutionnel", "Histoire"},
			Published:   true,
		},
		{
			Title:       "Simulation Architecturale",
			Slug:        "simulation-architecturale",
			Description: "Visualisation 3D photoréaliste d'un projet architectural pour présentation client.",
			Category:    "simulation_3d",
			Year:        2024,
			Client:      "Cabinet d'Architecture",
			CoverURL:    "",
			GalleryURLs: pq.StringArray{},
			Tags:        pq.StringArray{"Architecture", "3D", "Visualisation"},
			Published:   true,
		},
		{
			Title:       "Spot TV — Produit Local",
			Slug:        "spot-tv-produit-local",
			Description: "Production d'un spot télévisé pour la promotion d'un produit de consommation locale.",
			Category:    "spot_publicitaire",
			Year:        2024,
			Client:      "Marque Locale",
			CoverURL:    "",
			GalleryURLs: pq.StringArray{},
			Tags:        pq.StringArray{"Publicité", "TV", "Produit"},
			Published:   true,
		},
		{
			Title:       "Présentation de Projet Urbain",
			Slug:        "presentation-projet-urbain",
			Description: "Simulation 3D d'un projet d'aménagement urbain pour une présentation aux autorités.",
			Category:    "simulation_3d",
			Year:        2024,
			Client:      "Mairie",
			CoverURL:    "",
			GalleryURLs: pq.StringArray{},
			Tags:        pq.StringArray{"Urbanisme", "3D", "Présentation"},
			Published:   true,
		},
	}

	if err := db.Create(&items).Error; err != nil {
		log.Println("[seed] error seeding portfolio:", err)
		return
	}
	log.Println("[seed] portfolio seeded")
}

// ── Blog ──────────────────────────────────────────────────────────────────────

func seedBlog(db *gorm.DB) {
	var count int64
	db.Model(&models.Blog{}).Count(&count)
	if count > 0 {
		return
	}

	posts := []models.Blog{
		{
			Title:     "L'animation 3D en Afrique : état des lieux et perspectives",
			Slug:      "animation-3d-afrique",
			Excerpt:   "Comment l'animation 3D se développe-t-elle sur le continent africain ? Découvrez les tendances et opportunités.",
			Content:   "<p>L'animation 3D connaît un essor remarquable sur le continent africain. Des studios comme OPTIMAT Incorporation ouvrent la voie à une nouvelle génération de créateurs visuels.</p><h2>Un marché en pleine expansion</h2><p>La demande pour du contenu animé de qualité ne cesse de croître, portée par la multiplication des plateformes de streaming et la jeunesse de la population africaine.</p><h2>Les défis à relever</h2><p>Malgré cet élan, les studios africains font face à des défis importants : accès aux équipements, formation des talents, et financement des productions.</p><h2>OPTIMAT au cœur de cette révolution</h2><p>Fort de son expérience depuis 2016, OPTIMAT Incorporation s'impose comme un acteur incontournable de l'animation 3D en République Démocratique du Congo.</p>",
			Author:    "OPTIMAT Team",
			Category:  "Industrie",
			Published: true,
			CreatedAt: time.Date(2024, 11, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Title:     "Derrière les coulisses : la création de Koffi Raymet",
			Slug:      "coulisses-koffi-raymet",
			Excerpt:   "Plongez dans le processus créatif qui a donné naissance à la première série animée 3D congolaise.",
			Content:   "<p>Koffi Raymet est bien plus qu'une série animée : c'est le fruit de plusieurs années de travail acharné, d'innovations techniques et d'une passion sans limites pour la narration africaine.</p><h2>La genèse du projet</h2><p>L'idée est née en 2020, au sein des studios OPTIMAT. L'ambition : créer une série qui parle aux enfants congolais avec leurs propres héros, leurs propres décors.</p><h2>Les défis techniques</h2><p>La production d'une série animée 3D de qualité nécessite des ressources considérables. Rigger, texturer, animer chaque personnage a mobilisé toute l'équipe pendant deux ans.</p><h2>Un succès continental</h2><p>La saison 1 a été bien accueillie à travers l'Afrique centrale, ouvrant la voie à une saison 2 encore plus ambitieuse.</p>",
			Author:    "Yannick LFK",
			Category:  "Production",
			Published: true,
			CreatedAt: time.Date(2024, 10, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			Title:     "Les nouvelles technologies au service de l'animation",
			Slug:      "nouvelles-technologies-animation",
			Excerpt:   "Comment les avancées technologiques transforment la production d'animation 3D dans nos studios.",
			Content:   "<p>La révolution technologique touche de plein fouet l'industrie de l'animation. Chez OPTIMAT, nous avons fait le choix d'adopter les dernières innovations pour rester à la pointe de la création.</p><h2>L'intelligence artificielle dans le pipeline</h2><p>L'IA commence à jouer un rôle dans certaines étapes de la production : génération de textures, optimisation des rendus, assistance au rigging.</p><h2>Le rendu en temps réel</h2><p>Les moteurs de rendu en temps réel comme Unreal Engine 5 transforment notre manière de travailler, permettant des itérations plus rapides et des previsualisations instantanées.</p><h2>Notre engagement qualité</h2><p>Quelque soit la technologie utilisée, chez OPTIMAT, l'excellence artistique reste au cœur de chaque production.</p>",
			Author:    "OPTIMAT Team",
			Category:  "Technologie",
			Published: true,
			CreatedAt: time.Date(2024, 9, 20, 0, 0, 0, 0, time.UTC),
		},
	}

	if err := db.Create(&posts).Error; err != nil {
		log.Println("[seed] error seeding blog:", err)
		return
	}
	log.Println("[seed] blog seeded")
}
