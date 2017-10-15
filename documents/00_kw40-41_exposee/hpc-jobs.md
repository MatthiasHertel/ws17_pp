# Entwicklung einer REST API für HPC Jobs

# Beschreibung

Die API soll vor allem Serverseitig umgesetzt werden.

Der Transport/Kodierung der Daten erfolgt mittels HTTP und JSON.
Die Daten sollen Transport-verschlüsselt (TLS) übertragen werden.
Berechtigungen sollen von der API anhand geheimer Tokens geprüft werden.

Die API soll folgende Funktionen anbieten:

- Authentifizierung / Autorisierung mittels Tokens
- Upload der Inputfiles / Download der Outputfiles
- Anlegen eines Jobs mittels Job Templates und Verifikation der Meta-Daten
- Übergabe des Jobs an eine Queueing Engine, Abfrage des aktuellen Status und der Historie
- Der Job soll bei der Ausfuehrung Zugriff auf die Inputdaten bekommen und die Outputdaten anschliessend wieder in den Object-Store schreiben.
- Ergebnisse und Metadaten bzw. Ausgabe und Fehlermeldungen des Jobs werden zum Download angeboten

Insbesondere soll es möglich sein, dass ein Job über die beschriebene API neue Jobs in die Queue einstellen kann.
Dadurch soll der Upload von Variantensimulationen wesentlich vereinfacht und beschleunigt werden.

Es soll später viele Arten von Jobs geben und die Jobs sollen spezielle Anforderungen an die Laufzeit-Umgebung haben können (z.B. Anzahl Server und Cores pro Server).
Dies soll in der Entwicklung der REST API berücksichtigt werden. So soll es möglich sein, sich eine Liste aller möglichen Job-Templates und Laufzeit-Umgebungen ausgeben zu lassen.

Zugriff auf Ausgabe und Fehlermeldungen zur Laufzeit des Jobs (nicht erst nachdem er beendet ist) wäre wünschenswert, ist aber nicht zwingend notwendig.


## Umsetzung

Die API selbst soll auf existierender Software basieren. Eine Anlehnung an existierende APIs ist ausdrücklich gewünscht, aber nicht zwingend erforderlich.

Job-Templates sollen dabei aus Docker Containern und Adaptern zur Verifikation und Übergabe der Input bzw. Output-Daten bestehen.

Als Queueing System kann Univa Grid-Engine oder eine der vielen Open Source Alternativen zum Einsatz kommen.

Der Object-Storage soll kompatibel zu AWS S3 sein. RedHat Ceph oder Minio bieten sich an.

Ausserdem wird vermutlich eine Datenbank benötigt, die die Informationen aus den einzelnen Komponenten verknüpft. Zum aktuellen Zeitpunkt ist unklar, welche Art von Datenbank dafür am geeignetsten ist: relational, dokumentenorientiert oder Key-Value Store, verteilt oder nicht.

## Vorgehensweise

In iterativen Entwicklungszyklen von ca 2 Wochen sollen zuerst ein oder zwei Prototypen der REST Schnittstelle erstellt werden. Dies dient dazu, die geeignete Programmiersprache und vielleicht ein Framework zu finden.

Die verschiedenen Routes in der API werden per Mockup angelegt, bzw. dokumentiert. Möglicherweise bestehen einige Routen auch nur aus einem Proxy zum Object-Store.

Nebenläufig testen wir den Zugriff auf die Komponenten des Software Systems und überlegen, welche Anforderungen an die Datenbank-Modelle gestellt werden. Anschliessend soll die Integration gebaut werden, so dass für ein Beispiel-Jobtemplate der komplette Durchlauf funktioniert.
