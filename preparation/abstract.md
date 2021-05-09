Abstract
========

Die stark steigende Komplexität der Software Applikationen sowie die dauerhaft wachsende Anzahl an Nutzern zwingt den Softwaremarkt immer stärker in
die Cloud um, durch rapide Skalierung, mit dem Wachstum mithalten zu können. Monolithische Applikationen welche auf einzelnen bare-metal Maschinen
betrieben werden, können so kaum gegen die Microservices der Cloud antreten. Die dynamische Skalierung, welche die Cloud bietet, brachte auch ein
neues Kostenmodell mit sich welches stark auf den tatsächlich genutzten Ressourcen der Software, wie CPU, Arbeitsspeicher und Durchführungszeit
basiert.

Diese Arbeit diskutiert hierfür die Nutzung von GraalVM, einer stark auf Performance fokussierten Java Virtual Machine, sowie dem Quarkus Framework,
welches für die Nutzung innerhalb von Containers optimiert ist. Als erste Version mit Long-Term-Support bringt die, am 17. November 2020
veröffentlichte, 20.3.0 Version GraalVM in den Fokus des Enterprise Business. Das Ziel der Arbeit ist so der Vergleich von verschiedensten Java
Technologien zu der GraalVM/Quarkus Kombination. Hierbei wird die Arbeit vor allem auf die Frameworks Helidon, Micronaut und Spring sowie die OpenJDK
15 Virtual Machine eingehen, um ein möglichst genaues Abbild der momentanen Möglichkeiten zu bieten.

Auf Basis dieser Vergleiche wird diese Arbeit zusätzlich eine Kostenanalyse des Cloudhostings vornehmen um einen konkret, für die Enterprise
relevante, Aussage bezüglich der Effektivität der Technologien liefern zu können.



Notes
=====

Diese Arbeit wird als wissenschaftliche Nachfolge zu:
 - M. Šipek, D. Muharemagić, B. Mihaljević and A. Radovan, "Enhancing Performance of Cloud-based Software Applications with GraalVM and Quarkus," 2020
43rd International Convention on Information, Communication and Electronic Technology (MIPRO), Opatija, Croatia, 2020, pp. 1746-1751, doi:
10.23919/MIPRO48935.2020.9245290.
 - M. Šipek, B. Mihaljević and A. Radovan, "Exploring Aspects of Polyglot High-Performance Virtual Machine GraalVM,"
2019 42nd International Convention on Information and Communication Technology, Electronics and Microelectronics (MIPRO), Opatija, Croatia, 2019, pp.
1671-1676, doi: 10.23919/MIPRO.2019.8756917.

agieren.

Structure
=========

1. Einleitung
    1. Motivation
    2. Problemstellung
    3. Zielsetzung
    4. Forschungsmethodik
    5. Aufbau der Arbeit
2. Diskussion des theoretischen Hintergrunds
    1. Cloud Pricing Modell
    2. Java in Enterprise
    3. Quarkus Framework
    4. GraalVM
        1. GraalVM Native Images
        2. GraalVM und Quarkus
    5. Java Microservice
        1. Event Loop
        2. Synchrone / Reactive Entwicklung
3. Praktischer Vergleich der Microservices
    1. Zielsetzung und Forschungsmethodik
    2. Vorstellung des Microservice Clusters
    3. Analyse der Benchmarking Ergebnisse
4. Vergleich der wirtschaftlichen Auswirkungen
    1. Migrationskosten
    2. Prozesskosten
    3. Speicherkosten
5. Kritische Reflexion und Ausblick

Sources
=======

- Laatikainen, G., Ojala, A., & Mazhelis, O. (2013). Cloud services pricing models. Lecture Notes in Business Information Processing, 150 LNBIP, 117–129. https://doi.org/10.1007/978-3-642-39336-5_12
- Albert, E., Arenas, P., Genaim, S., Puebla, G., & Zanardini, D. (2007). Cost analysis of Java bytecode. Lecture Notes in Computer Science (Including Subseries Lecture Notes in Artificial Intelligence and Lecture Notes in Bioinformatics), 4421 LNCS, 157–172. https://doi.org/10.1007/978-3-540-71316-6_12
- Xu, H., & Li, B. (2013). Dynamic Cloud Pricing for Revenue Maximization. IEEE Transactions on Cloud Computing, 1(2), 158–171. https://doi.org/10.1109/TCC.2013.15
- Sipek, M., Muharemagic, D., Mihaljevic, B., & Radovan, A. (2020). Enhancing performance of cloud-based software applications with GraalVM and quarkus. 2020 43rd International Convention on Information, Communication and Electronic Technology, MIPRO 2020 - Proceedings, 1746–1751. https://doi.org/10.23919/MIPRO48935.2020.9245290
- Evaluation of GraalVM Performance for Java Programs. (n.d.).
- Šipek, M., Mihaljevic, B., & Radovan, A. (2019). Exploring aspects of polyglot high-performance virtual machine GraalVM. 2019 42nd International Convention on Information and Communication Technology, Electronics and Microelectronics, MIPRO 2019 - Proceedings, 1671–1676. https://doi.org/10.23919/MIPRO.2019.8756917
- Al-Haidari, F., Sqalli, M., & Salah, K. (2013). Impact of CPU utilization thresholds and scaling size on autoscaling cloud resources. Proceedings of the International Conference on Cloud Computing Technology and Science, CloudCom, 2, 256–261. https://doi.org/10.1109/CloudCom.2013.142
- Ankit Kumar Sah. (2013). JAVA WEB DEPLOYMENT IN CLOUD COMPUTING. http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.402.7230&rep=rep1&type=pdf
- Maenhaut, P.-J., Moens, H., Ongenae, V., & De Turck, F. (2016). Migrating legacy software to the cloud: approach and verification by means of two medical software use cases. Software: Practice and Experience, 46(1), 31–54. https://doi.org/10.1002/spe.2320
- Würthinger, T., Wimmer, C., Wöß, A., Stadler, L., Duboscq, G., Humer, C., Richards, G., Simon, D., & Wolczko, M. (n.d.). One VM to Rule Them All. https://doi.org/10.1145/2509578.2509581
- Kash, I. A., & Key, P. B. (2016). Pricing the cloud. IEEE Internet Computing, 20(1), 36–43. https://doi.org/10.1109/MIC.2016.4
- Niephaus, F., Felgentreff, T., & Hirschfeld, R. (2019). Towards polyglot adapters for the GraalVM. ACM International Conference Proceeding Series, 19, 1–3. https://doi.org/10.1145/3328433.3328458
- Simão, J., & Veiga, L. (2012). VM economics for Java cloud computing: An adaptive and resource-aware Java runtime with quality-of-execution. Proceedings - 12th IEEE/ACM International Symposium on Cluster, Cloud and Grid Computing, CCGrid 2012, 723–728. https://doi.org/10.1109/CCGrid.2012.121
- Wimmer, C., Stancu, C., Hofer, P., Jovanovic, V., Wögerer, P., Kessler, P. B., Pliss, O., & Würthinger, T. (2019). Initialize once, start fast: application initialization at build time. Proceedings of the ACM on Programming Languages, 3(OOPSLA), 1–29. https://doi.org/10.1145/3360610
- Huijie Pan. (2020). INTERFACE LIVE GRAPH WITH RAPIDWRIGHT [UNIVERSITY OF CALIFORNIA]. https://escholarship.org/content/qt57b4s6nf/qt57b4s6nf_noSplash_0b7446d0f5f76c2736bb388b4bd4a0b1.pdf
- Björk, K. (2020). A comparison of compiler strategies for serverless functions written in Kotlin. In DEGREE PROJECT COMPUTER SCIENCE AND ENGINEERING. http://urn.kb.se/resolve?urn=urn:nbn:se:kth:diva-273961


