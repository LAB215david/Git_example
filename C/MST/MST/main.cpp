// C++ code
//參考--http://alrightchiu.github.io/SecondRound/minimum-spanning-treekruskals-algorithm.html
#include <iostream>
#include <vector>
#include <list>
#include <iomanip>      // for setw()

//每條線的架構
struct Edge{
    int from, to, weight;
    Edge(){};
    Edge(int u, int v, int w):from(u), to(v), weight(w){};
};

//
class GraphMST{
private:
    int num_vertex;
    std::vector<std::vector<int>> AdjMatrix;
public:
    GraphMST():num_vertex(0){};
    GraphMST(int n):num_vertex(n){
        AdjMatrix.resize(num_vertex);
        for (int i = 0; i < num_vertex; i++) {
            AdjMatrix[i].resize(num_vertex);
        }
    }
    void AddEdge(int from, int to, int weight);

    void KruskalMST();
    void GetSortedEdge(std::vector<struct Edge> &vec);
    friend int FindSetCollapsing(int *subset, int i);
    friend void UnionSet(int *subset, int x, int y);
};

// 用遞迴做collapsing
// subset：用來記錄edgesetMST[]中的edge之兩端vertex所屬的集合，目的是用來判斷是否形成cycle
int FindSetCollapsing(int *subset, int i){      

    int root;  // root
    for (root = i; subset[root] >= 0; root = subset[root]);

    while (i != root) {
        int parent = subset[i];
        subset[i] = root;
        i = parent;
    }

    return root;
}

//
void UnionSet(int *subset, int x, int y){

    int xroot = FindSetCollapsing(subset, x),
        yroot = FindSetCollapsing(subset, y);

    // 用rank比較, 負越多表示set越多element, 所以是值比較小的element比較多
    // xroot, yroot的subset[]一定都是負值
    if (subset[xroot] <= subset[yroot]) {        // x比較多element或是一樣多的時候, 以x作為root
        subset[xroot] += subset[yroot];
        subset[yroot] = xroot;
    }
    else {    //  if (subset[xroot] > subset[yroot]), 表示y比較多element
        subset[yroot] += subset[xroot];
        subset[xroot] = yroot;
    }
}

//
bool WeightComp(struct Edge e1, struct Edge e2){
    return (e1.weight < e2.weight);
}

//
void GraphMST::GetSortedEdge(std::vector<struct Edge> &edgearray){

    for (int i = 0; i < num_vertex-1; i++) {
        for (int j = i+1; j < num_vertex; j++) {
            if (AdjMatrix[i][j] != 0) {
                edgearray.push_back(Edge(i,j,AdjMatrix[i][j]));
            }
        }
    }
    // 用std::sort 排序, 自己定義一個comparison
    std::sort(edgearray.begin(), edgearray.end(), WeightComp);
}

//
void GraphMST::KruskalMST(){
    //edgesetMST用來收集所有MST中的edge
    struct Edge *edgesetMST = new struct Edge[num_vertex-1];
    int edgesetcount = 0;

    int subset[num_vertex];
    for (int i = 0; i < num_vertex; i++) {
        subset[i] = -1;
    }

// increaseWeight：把Graph中的edge按照weight由小到大排序，依序放進increaseWeight[]，當演算法在「挑選edge」形成最短路徑時，便是按照「weight由小到大」之順序挑選。
    std::vector<struct Edge> increaseWeight;
    GetSortedEdge(increaseWeight);              // 得到排好序的edge的vec

    for (int i = 0; i < increaseWeight.size(); i++) {
        if (FindSetCollapsing(subset, increaseWeight[i].from) != FindSetCollapsing(subset, increaseWeight[i].to)) {
            edgesetMST[edgesetcount++] = increaseWeight[i];
            UnionSet(subset, increaseWeight[i].from, increaseWeight[i].to);
        }
    }
    // 印出vertex與vertex之predecessor
    std::cout << std::setw(3) << "v1" << " - " << std::setw(3) << "v2"<< " : weight\n";
    for (int i = 0; i < num_vertex-1; i++) {
        std::cout << std::setw(3) << edgesetMST[i].from << " - " << std::setw(3) << edgesetMST[i].to 
                  << " : " << std::setw(4) << edgesetMST[i].weight << "\n";
    }
}

//將路線加入matrix
void GraphMST::AddEdge(int from, int to, int weight){
    AdjMatrix[from][to] = weight;
}

int main(){

    GraphMST g10(10);
    g10.AddEdge(0, 1, 2);g10.AddEdge(0, 2, 1);g10.AddEdge(0, 3, 4);
    g10.AddEdge(1, 0, 2);g10.AddEdge(1, 2, 2);g10.AddEdge(1, 4, 5);
    g10.AddEdge(2, 0, 1);g10.AddEdge(2, 1, 2);g10.AddEdge(2, 3, 4);g10.AddEdge(2, 4, 3);
    g10.AddEdge(3, 0, 4);g10.AddEdge(3, 2, 4);g10.AddEdge(3, 4, 6);g10.AddEdge(3, 5, 6);g10.AddEdge(3, 7, 2);
    g10.AddEdge(4, 1, 5);g10.AddEdge(4, 2, 3);g10.AddEdge(4, 3, 6);g10.AddEdge(4, 5, 3);g10.AddEdge(4, 6, 2);g10.AddEdge(4, 8, 4);
    g10.AddEdge(5, 3, 6);g10.AddEdge(5, 4, 3);g10.AddEdge(5, 7, 5);g10.AddEdge(5, 8, 8);
    g10.AddEdge(6, 4, 2);g10.AddEdge(6, 8, 7);
    g10.AddEdge(7, 3, 2);g10.AddEdge(7, 5, 5);g10.AddEdge(7, 8, 8);g10.AddEdge(7, 9, 9);
    g10.AddEdge(8, 4, 4);g10.AddEdge(8, 5, 8);g10.AddEdge(8, 6, 7);g10.AddEdge(8, 7, 8);g10.AddEdge(8, 9, 1);
    g10.AddEdge(9, 7, 9);g10.AddEdge(9, 8, 1);
    
    std::cout << "MST found by Kruskal:\n";
    g10.KruskalMST();

    return 0;  
}
