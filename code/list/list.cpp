#include "stdafx.h"
#include<iostream>
#include "List.h"
using namespace std;

Status List_Retrieve(ListPtr L, int pos, ElemType * elem) {
    Status status = rangeerror;
    int len = L->length;
    if (1 <= pos&&pos <= len) {
        *elem = L->elem[pos];
        status = success;
    }
    return status;
}

Status List_Locate(ListPtr L, ElemType elem, int * pos) {
    Status status = rangeerror;
    int len = L->length;
    int i = 1;
    while (i <= len && L->elem[i]==elem) {
        i++;
    }
    if (i <= len) {
        *pos = i;
        status = success;
    }
    return status;
}

Status List_Insert(ListPtr L, int pos, ElemType elem) {
    Status status = rangeerror;
    int len = L->length,i;
    if (len > MAXSIZE)status = overflow;
    else if (1 <= pos && pos <= len + 1) {
        for (i = len; i >= pos; i--)
            L->elem[i + 1] = elem;
        L->elem[pos] = elem;
        L->length++;
        status = success;
    }
    return status;
}

Status List_Init(ListPtr L) {
    Status status = fatal;
    L->elem = (ElemType*)malloc((MAXSIZE + 1) * sizeof(ElemType));
    if (L->elem) {
        L->length = 0;
        status = success;
    }
    return status;
}

void List_Destroy(ListPtr L) {
    if (L->elem) {
        free(L->elem);
        L->elem = NULL;
    }
    L->length = 0;
}

void List_Clear(ListPtr L) {
    L->length = 0;
}

bool List_Empty(ListPtr L) {
    return L->length == 0;
}

Status List_Prior(ListPtr L, int pos, ElemType * elem) {
    Status status = fail;
    int len = L->length;
    if (2 <= pos && pos <= len) {
        *elem = L->elem[pos - 1];
        status = success;
    }
    return status;
}

Status List_Next(ListPtr L, int pos, ElemType * elem) {
    Status status;
    int len = L->length;
    if (1 <= pos && pos <= len - 1) {
        *elem = L->elem[pos+1];
        status = success;
    }
    return status;
}

int List_Size(ListPtr L) {
    return L->length;
}

Status List_Remove(ListPtr L, int pos) {
    Status status = rangeerror;
    int len = L->length, i;
    if (1 <= pos && pos <= len) {
        for (i = pos; i < len; i++) 
            L->elem[i] = L->elem[i + 1];
        L->length--;
        status = success;
    }
    return status;
}


Status List_Union(ListPtr La, ListPtr Lb) {
    ElemType elem;
    Status status= fail;
    int i, len = List_Size(Lb);
    for (i = 1; i <= len; i++) {
        List_Retrieve(Lb, i, &elem);
        if (status != success) {
            status = List_Insert(La, 1, elem);
            if (status != success)break;
        }
    }
    return status;
}



Status List_Merge(ListPtr La, ListPtr Lb, ListPtr Lc) {
    ElemType elem1, elem2;
    Status status=fail;
    status = List_Init(Lc);
    if(status != success)return status;
    int i = 1, j = 1, k = 1;
    int n = List_Size(La), m = List_Size(Lb);
    while (i <= n && j <= m) {
        List_Retrieve(La, i, &elem1), List_Retrieve(Lb, j, &elem2);
        if (elem1 < elem2) {
            status = List_Insert(Lc, k, elem1);
            i++;
        }
        else {
            status = List_Insert(Lc, k, elem2);
            j++;
        }
        if (status != success) return status;
        k++;
    }
    while (i <= n) {
        List_Retrieve(La, i, &elem1);
        status = List_Insert(Lc, k, elem1);
        if (status != success) return status;
        i++, k++;
    }
    while (j <= m) {
        List_Retrieve(Lb, j, &elem2);
        status = List_Insert(Lc, k, elem2);
        if (status != success) return status;
        j++, k++;
    }
    return status;
}


int main() {
    ListPtr La = new Sqlist, Lb = new Sqlist, Lc = new Sqlist;;
    List_Init(La);
    List_Init(Lb);
    List_Init(Lc);
    int arra[5] = { 2,4,6,7,8 };
    for (int i = 1; i <= 5; i++)
        List_Insert(La, i, arra[i-1]);
    
    for(int i=1;i<=5;i++)
        cout << La->elem[i] << " ";
    cout << endl;

    int arrb[5] = { 1,5,9,10,11 };
    for (int i = 1; i <= 5; i++)
        List_Insert(Lb, i, arrb[i - 1]);

    for (int i = 1; i <= 5; i++)
        cout << Lb->elem[i] << " ";
    cout<< endl;

    Status status = List_Merge(La, Lb, Lc);
    cout << status << endl;
    if (status != success)return EXIT_FAILURE;

    for (int i = 1; i <= 10; i++)
        cout << Lc->elem[i] << " ";
    cout << endl;


    system("pause");
    return EXIT_SUCCESS;
}