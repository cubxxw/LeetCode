#pragma once
#include "targetver.h"

#include <stdio.h>
#include <tchar.h>

#define LIST_INIT_SIZE 100
#define MAXSIZE 100

typedef int ElemType;

typedef struct Sqlist {
    ElemType * elem;
    int length;
    int listsize;
}Sqlist, *ListPtr;

typedef enum Status {
    success, fail, fatal, rangeerror,overflow
}Status;

Status List_Init(ListPtr L);

void List_Destroy(ListPtr L);

void List_Clear(ListPtr L);

bool List_Empty(ListPtr L);

int List_Size(ListPtr L);

Status List_Retrieve(ListPtr L, int pos, ElemType * elem);

Status List_Locate(ListPtr L, ElemType elem, int *pos);

Status List_Insert(ListPtr L, int pos, ElemType elem);

Status List_Remove(ListPtr L, int pos);

Status List_Prior(ListPtr L, int pos, ElemType * elem);