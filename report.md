### Machine Learning - The k-means algorithm


Implementing the K-means clustering algorithm on the 4 dimensional, 3 dimensional csv dataset:


read dataset's with the read_csv method :

``
    def read_csv():
        x = []
        y = []
        labels = []
        x_label = ""
        y_label = ""
        with open('../csc/aggregation.csv') as csvfile:
            reader = csv.reader(csvfile, delimiter=',')
            lines = 0

            #iterate through each row in aggregation.csv and append to it's respective label i.e list
            for row in reader:
                if lines >= 1:
                    print(', '.join(row))
                    x.append(float(row[1]))
                    y.append(float(row[2]))
                    labels.append(row[0])
                    lines += 1
                else:
                    x_label = row[1]
                    y_label = row[2]
                    print(', '.join(row))
                    lines += 1
    return x, y, x_label, y_label, labels
``

unpack csv data into x, y list of data by calling the read_csv() method
`x, y, x_label, y_label, labels = read_csv()`


combine x, y into a 2 dimensional lists of (x, y) pairs using numpy

`X = np.vstack((x, y)).T`

finding the k-mean cluster

`find_clusters(X, n_clusters, rseed=2):`

the find cluster function randomly picks centres in our 2D list of x, y pairs i.e `X`

using np.random.RandomState(rseed)

```
rng = np.random.RandomState(rseed)
i = rng.permutation(X.shape[0])[:n_clusters]
centers = X[i]
```

Next we  iterate through the X numpy array for new means and check for convergence

``
    new_centers = np.array([X[labels == i].mean(0) for i in
    range(n_clusters)])

    # 2c. Check for convergence
    if np.all(centers == new_centers):
        break
    centers = new_centers
``
once theres a convergence we break out of the loop and return our centers 
![k-mean centres](centres.png)




If we further use `matplotlib` to visualize the dataset we get a pictorial representation of the k-means clusters:

![aggregates](Tasks/aggregates.png)




next we find the the number of `*k*` at the elbow of the plot

`distortions = []`
average of the squared distances from the cluster centers of the respective clusters

`inertias = [] `
sum of squared differences from the closest mean


iterating using 2 through 11 means, `K = range(2, 11)`
``
for k in K:
    # Building and fitting the model
    kmeanModel = KMeans(n_clusters=k).fit(X)
    kmeanModel.fit(X)
 
    distortions.append(sum(np.min(cdist(X, kmeanModel.cluster_centers_,
                                        'euclidean'), axis=1)) / X.shape[0])
    inertias.append(kmeanModel.inertia_)
 
    mapping1[k] = sum(np.min(cdist(X, kmeanModel.cluster_centers_,
                                   'euclidean'), axis=1)) / X.shape[0]
    mapping2[k] = kmeanModel.inertia_
``


We can now visualize our `*k*` elbow of the plots


![k elbow plots](elbows.png)

``
for key, val in mapping1.items():
    print(f'{key} : {val}')
``

![k elbow plots](dialog.png)

