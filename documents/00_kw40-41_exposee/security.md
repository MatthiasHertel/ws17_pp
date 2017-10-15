# Entwicklung einer Security Suite


## Motivation

Als Administratoren von HPC Clustern haben wir das Problem, dass wir auf bekannt gewordene Sicherheitslücken in der eingesetzten Software reagieren müssen. Dies kann aber nicht immer durch ein zeitnahes Update geschehen, denn die bei uns gerechneten HPC Jobs können so gross sein, dass sie auf mehreren Servern mehrere Wochen rechnen. Diese Server kann man in dieser Zeit nicht updaten.

## Beschreibung

Um dem Kunden das Problem bewusst zu machen, benötigen wir einen Abgleich, welche Security Probleme auf welchem Server, Docker Container, Virtuellen Maschine oder Firewall vorhanden sind. Um das eingrenzen zu koennen, soll die Software-Version laut Paketverwaltung (es gibt auf einem Server potentiell mehrere Paketverwaltungen) vorhanden sind.

Der Versionsstand von Software, die nicht über Paketverwaltungen installiert wird, soll möglichst einfach im Terminal eingepflegt werden können. Als Beispiel soll z.B. nach einem Git Checkout der aktuelle Commit Hash (plus Hostname, Installationsverzeichnis und Zeitstempel) erfasst werden.

Diese Informationen sollen an einer zentralen Stelle im Cluster hinterlegt werden. Die Zusammenfassung der Installation über Cluster-Grenzen hinweg soll möglich sein.

Ausserdem wären historische Daten (insbesondere Änderungen) wichtig.

### Software
